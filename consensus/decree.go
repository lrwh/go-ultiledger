package consensus

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/deckarep/golang-set"
	pb "github.com/golang/protobuf/proto"

	"github.com/ultiledger/go-ultiledger/log"
	"github.com/ultiledger/go-ultiledger/ultpb"
)

type BallotPhase uint8

const (
	BallotPhasePrepare BallotPhase = iota
	BallotPhaseConfirm
	BallotPhaseExternalize
)

// Decree is an abstractive decision the consensus engine
// should reach in each round
type Decree struct {
	index           uint64
	nodeID          string
	quorum          *ultpb.Quorum
	quorumHash      string
	latestCloseTime uint64

	// for nomination protocol
	votes            mapset.Set
	accepts          mapset.Set
	candidates       mapset.Set
	nominations      map[string]*ultpb.Nomination
	nominationRound  int
	latestNomination *ultpb.Nomination
	latestComposite  string // latest composite candidate value

	// for ballot protocol
	currentPhase     BallotPhase
	currentBallot    *ultpb.Ballot
	pBallot          *ultpb.Ballot // p
	qBallot          *ultpb.Ballot // p'
	hBallot          *ultpb.Ballot // h
	cBallot          *ultpb.Ballot // c
	ballots          map[string]*ultpb.Statement
	latestBallotStmt *ultpb.Statement

	// channel for sending statements
	statementChan chan *ultpb.Statement
}

func NewDecree(idx uint64, nodeID string, quorum *ultpb.Quorum, quorumHash string, stmtC chan *ultpb.Statement) *Decree {
	d := &Decree{
		index:           idx,
		nodeID:          nodeID,
		quorum:          quorum,
		quorumHash:      quorumHash,
		nominationRound: 0,
		votes:           mapset.NewSet(),
		accepts:         mapset.NewSet(),
		candidates:      mapset.NewSet(),
		nominations:     make(map[string]*ultpb.Nomination),
		statementChan:   stmtC,
	}
	return d
}

// Nominate nominates a consensus value for the decree
func (d *Decree) Nominate(prevHash, currHash string) error {
	d.nominationRound++
	// TODO(bobonovski) compute leader weights
	d.votes.Add(currHash) // For test

	if err := d.sendNomination(); err != nil {
		return fmt.Errorf("send nomination failed: %v", err)
	}
	return nil
}

// Recv receives validated statement and redistributes it to
// corresponding route handler. If the statement is a nomination,
// we give it to seperated nomination handler. Other statement types
// belong to ballot protocol, we directly pass the statement to the
// handler which contains its own logic to distinguish fine grained
// statement types.
func (d *Decree) Recv(stmt *ultpb.Statement) error {
	if stmt == nil {
		return errors.New("statement is nil")
	}

	switch stmt.StatementType {
	case ultpb.StatementType_NOMINATE:
		nom, err := ultpb.DecodeNomination(stmt.Data)
		if err != nil {
			return fmt.Errorf("decode nomination failed: %v", err)
		}
		err = d.recvNomination(d.nodeID, nom)
		if err != nil {
			return fmt.Errorf("recv nomination failed: %v", err)
		}
	case ultpb.StatementType_PREPARE:
		fallthrough
	case ultpb.StatementType_CONFIRM:
		fallthrough
	case ultpb.StatementType_EXTERNALIZE:
		err := d.recvBallot(stmt)
		if err != nil {
			return fmt.Errorf("recv ballot failed: %v", err)
		}
	}

	return nil
}

/* Nomination Protocol */
// receive nomination from peers or local node
func (d *Decree) recvNomination(nodeID string, nom *ultpb.Nomination) error {
	// check validity of votes and accepts
	if len(nom.VoteList)+len(nom.AcceptList) == 0 {
		return errors.New("vote and accept list is empty")
	}

	// check whether the existing nomination of the remote node
	// is the proper subset of the new nomination
	if oldNom, ok := d.nominations[nodeID]; ok {
		if isNewerNomination(oldNom, nom) {
			d.nominations[nodeID] = nom
		}
	}
	acceptUpdated, candidateUpdated, err := d.promoteVotes(nom)
	if err != nil {
		return fmt.Errorf("promote votes failed: %v", err)
	}

	// send new nomination if votes changed
	if acceptUpdated {
		d.sendNomination()
	}

	// start balloting if candidates changed
	if candidateUpdated {
		compValue, err := d.combineCandidates()
		if err != nil {
			return fmt.Errorf("combine candidates failed: %v", err)
		}
		d.latestComposite = compValue

		d.updateBallotPhase(compValue, false)
	}

	return nil
}

// assemble a nomination and broadcast it to other peers
func (d *Decree) sendNomination() error {
	// create an abstract nomination statement
	nom := &ultpb.Nomination{
		QuorumHash: d.quorumHash,
	}
	for vote := range d.votes.Iter() {
		nom.VoteList = append(nom.VoteList, vote.(string))
	}
	for accept := range d.accepts.Iter() {
		nom.AcceptList = append(nom.AcceptList, accept.(string))
	}

	if err := d.recvNomination(d.nodeID, nom); err != nil {
		return fmt.Errorf("receive local nomination failed: %v", err)
	}

	// broadcast the nomination if it is a new one
	if isNewerNomination(d.latestNomination, nom) {
		d.latestNomination = nom
		nomBytes, err := ultpb.Encode(nom)
		if err != nil {
			return fmt.Errorf("encode nomination failed: %v", err)
		}
		stmt := &ultpb.Statement{
			StatementType: ultpb.StatementType_NOMINATE,
			NodeID:        d.nodeID,
			Index:         d.index,
			Data:          nomBytes,
		}
		d.statementChan <- stmt
	}

	return nil
}

// try to promote votes to accepts by checking two conditions (ACCEPT):
//   1. whether the votes form V-blocking
//   2. whether all the nodes in the quorum have voted
// then try to promote accepts to candidates by checking (CONFIRM):
//   1. whether all the nodes in the quorum have accepted
func (d *Decree) promoteVotes(newNom *ultpb.Nomination) (bool, bool, error) {
	acceptUpdated := false
	for _, vote := range newNom.VoteList {
		if d.accepts.Contains(vote) {
			continue
		}

		// use federated vote to promote value
		ns := findAcceptNodes(vote, d.nominations)
		if !isVblocking(d.quorum, ns) {
			nset := findVoteOrAcceptNodes(vote, d.nominations)
			if !isQuorumSlice(d.quorum, nset) { // TODO(bobonovski) trim nset to contain only other quorums
				return false, false, fmt.Errorf("failed to promote any votes to accepts")
			}
		}

		// TODO(bobonovski) check the validity of the vote
		d.votes.Add(vote)
		d.accepts.Add(vote)
		acceptUpdated = true
	}

	candidateUpdated := false
	for _, accept := range newNom.AcceptList {
		if d.candidates.Contains(accept) {
			continue
		}

		ns := findAcceptNodes(accept, d.nominations)
		if isQuorumSlice(d.quorum, ns) {
			d.candidates.Add(accept)
			candidateUpdated = true
		}
	}

	return acceptUpdated, candidateUpdated, nil
}

func (d *Decree) combineCandidates() (string, error) {
	return "", nil
}

/* Ballot Protocol */
// receive ballot statement from peer or local nodes
func (d *Decree) recvBallot(stmt *ultpb.Statement) error {
	if stmt.Index != d.index {
		log.Fatalf("received incompatible ballot index: local %d, recv %d", d.index, stmt.Index)
	}

	// skip outdated statement without returning error
	if s, ok := d.ballots[stmt.NodeID]; ok {
		if !isNewerBallot(s, stmt) {
			return nil
		}
	}

	// make sure the ballot is valid
	if err := d.validateBallot(stmt); err != nil {
		return fmt.Errorf("ballot is invalid: %v", err)
	}

	if d.currentPhase != BallotPhaseExternalize {
		d.ballots[stmt.NodeID] = stmt
	} else {

	}

	return nil
}

// assemble a ballot statement based on current ballot phase
func (d *Decree) sendBallot() error {
	var stmtType ultpb.StatementType
	var msg pb.Message

	d.checkBallotInvariants()

	// assemble ballot statement based on current phase
	switch d.currentPhase {
	case BallotPhasePrepare: // Prepare statement
		stmtType = ultpb.StatementType_PREPARE
		prepare := &ultpb.Prepare{
			B:          d.currentBallot,
			P:          d.pBallot,
			Q:          d.qBallot,
			QuorumHash: d.quorumHash,
		}
		if d.cBallot != nil {
			prepare.LC = d.cBallot.Counter
		}
		if d.hBallot != nil {
			prepare.HC = d.hBallot.Counter
		}
		msg = prepare
	case BallotPhaseConfirm: // Confirm statement
		stmtType = ultpb.StatementType_CONFIRM
		confirm := &ultpb.Confirm{
			B:          d.currentBallot,
			PC:         d.pBallot.Counter,
			LC:         d.cBallot.Counter,
			HC:         d.hBallot.Counter,
			QuorumHash: d.quorumHash,
		}
		msg = confirm
	case BallotPhaseExternalize: // Externalize statement
		stmtType = ultpb.StatementType_EXTERNALIZE
		exter := &ultpb.Externalize{
			B:          d.cBallot,
			HC:         d.hBallot.Counter,
			QuorumHash: d.quorumHash,
		}
		msg = exter
	default:
		log.Fatalf("invalid ballot phase: %d", d.currentPhase)
	}

	// create statement
	msgBytes, err := ultpb.Encode(msg)
	if err != nil {
		return fmt.Errorf("encode ballot failed: %v", err)
	}
	stmt := &ultpb.Statement{
		StatementType: stmtType,
		NodeID:        d.nodeID,
		Index:         d.index,
		Data:          msgBytes,
	}

	// check whether the statement is already processed
	s, ok := d.ballots[d.nodeID]
	if !ok || pb.Equal(s, stmt) {
		if err := d.recvBallot(stmt); err != nil {
			return fmt.Errorf("recv local ballot failed: %v", err)
		}
		if d.latestBallotStmt == nil || isNewerBallot(d.latestBallotStmt, stmt) {
			d.latestBallotStmt = stmt
			// broadcast the ballot
			d.statementChan <- stmt
		}
	}

	return nil
}

// try to step ballot state
func (d *Decree) step(stmt *ultpb.Statement) error {
	return nil
}

// try to accept new ballot as prepared
func (d *Decree) acceptPrepared(stmt *ultpb.Statement) error {
	// it is only necessary to call this method when
	// current phase is in prepare or confirm.
	if d.currentPhase != BallotPhasePrepare && d.currentPhase != BallotPhaseConfirm {
		return fmt.Errorf("current phase not in prepare or confirm: %d", d.currentPhase)
	}
	return nil
}

// extract unique prepare candidate ballots from statement
func (d *Decree) extractPrepareCandidates(stmt *ultpb.Statement) ([]*ultpb.Ballot, error) {
	// filter ballots with the same value
	ballots := mapset.NewSet()

	switch stmt.StatementType {
	case ultpb.StatementType_PREPARE:
		prepare, err := ultpb.DecodePrepare(stmt.Data)
		if err != nil {
			return nil, fmt.Errorf("decode prepare statement failed: %v", err)
		}
		ballots.Add(*prepare.B)
		if prepare.P != nil {
			ballots.Add(*prepare.P)
		}
		if prepare.Q != nil {
			ballots.Add(*prepare.Q)
		}
	case ultpb.StatementType_CONFIRM:
		confirm, err := ultpb.DecodeConfirm(stmt.Data)
		if err != nil {
			return nil, fmt.Errorf("decode confirm statement failed: %v", err)
		}
		ballots.Add(ultpb.Ballot{Value: confirm.B.Value, Counter: confirm.PC})
		ballots.Add(ultpb.Ballot{Value: confirm.B.Value, Counter: math.MaxUint32})
	case ultpb.StatementType_EXTERNALIZE:
		ext, err := ultpb.DecodeExternalize(stmt.Data)
		if err != nil {
			return nil, fmt.Errorf("decode externalize statement failed: %v", err)
		}
		ballots.Add(ultpb.Ballot{Value: ext.B.Value, Counter: math.MaxUint32})
	default:
		log.Fatalf("invalid ballot statement type: %d", stmt.StatementType)
	}

	var candidates []*ultpb.Ballot
	if ballots.Cardinality() == 0 {
		return candidates, nil
	}

	// process ballots in descending order
	candSet := mapset.NewSet()
	for ballot := range ballots.Iter() {
		b := ballot.(ultpb.Ballot)
		for _, stmt := range d.ballots {
			switch stmt.StatementType {
			case ultpb.StatementType_PREPARE:
				prepare, err := ultpb.DecodePrepare(stmt.Data)
				if err != nil {
					// skip corrupted ballot
					continue
				}
				if lessAndCompatibleBallots(prepare.B, &b) {
					candSet.Add(b)
				}
				if prepare.P != nil && lessAndCompatibleBallots(prepare.P, &b) {
					candSet.Add(*prepare.P)
				}
				if prepare.Q != nil && lessAndCompatibleBallots(prepare.Q, &b) {
					candSet.Add(*prepare.Q)
				}
			case ultpb.StatementType_CONFIRM:
				confirm, err := ultpb.DecodeConfirm(stmt.Data)
				if err != nil {
					// skip corrupted ballot
					continue
				}
				if compatibleBallots(confirm.B, &b) {
					candSet.Add(b)
					if confirm.PC < b.Counter {
						candSet.Add(ultpb.Ballot{Value: b.Value, Counter: confirm.PC})
					}
				}
			case ultpb.StatementType_EXTERNALIZE:
				ext, err := ultpb.DecodeExternalize(stmt.Data)
				if err != nil {
					// skip corrupted ballot
					continue
				}
				if compatibleBallots(ext.B, &b) {
					candSet.Add(b)
				}
			default:
				log.Fatalf("invalid ballot statement type: %d", stmt.StatementType)
			}
		}
	}

	for v := range candSet.Iter() {
		b := v.(ultpb.Ballot)
		candidates = append(candidates, &b)
	}

	return candidates, nil
}

// update the current ballot phase
func (d *Decree) updateBallotPhase(val string, force bool) bool {
	if !force && d.currentBallot == nil {
		return false
	}

	counter := uint32(1)
	if d.currentBallot != nil {
		counter = d.currentBallot.Counter + 1
	}

	if d.currentPhase != BallotPhasePrepare && d.currentPhase != BallotPhaseConfirm {
		return false
	}

	// TODO(bobonovski) use confirmed prepared value?
	b := &ultpb.Ballot{Counter: counter, Value: val}

	updated := d.updateBallotValue(b)
	if updated {
		if err := d.sendBallot(); err != nil {
			log.Fatalf("send ballot failed: %v", err)
		}
	}

	return updated
}

// update the current ballot value
func (d *Decree) updateBallotValue(b *ultpb.Ballot) bool {
	if d.currentPhase != BallotPhasePrepare && d.currentPhase != BallotPhaseConfirm {
		return false
	}

	updated := false

	if d.currentBallot == nil {
		d.updateBallot(b)
		updated = true
	} else {
		if compareBallots(d.currentBallot, b) <= 0 {
			log.Fatal("cannot update current ballot with smaller one")
		}

		if d.cBallot != nil && strings.Compare(d.cBallot.Value, b.Value) != 0 {
			return false
		}

		if compareBallots(d.currentBallot, b) <= 0 {
			d.updateBallot(b)
			updated = true
		}
	}

	d.checkBallotInvariants()

	return updated
}

// update the current ballot
func (d *Decree) updateBallot(b *ultpb.Ballot) {
	if d.currentPhase == BallotPhaseExternalize {
		log.Fatal("should not update ballot in externalize phase")
	}

	if d.currentBallot != nil && compareBallots(d.currentBallot, b) <= 0 {
		log.Fatal("cannot update current ballot with smaller one")
	}

	d.currentBallot = &ultpb.Ballot{Counter: b.Counter, Value: b.Value}

	if d.hBallot != nil && !compatibleBallots(d.currentBallot, d.hBallot) {
		d.hBallot.Reset()
	}
}

// check invariants of ballot states
func (d *Decree) checkBallotInvariants() {
	if d.currentBallot != nil && d.currentBallot.Counter == 0 {
		log.Fatal("current ballot is not nil but counter is zero")
	}

	if d.pBallot != nil && d.qBallot != nil {
		cond := compareBallots(d.qBallot, d.pBallot) <= 0 && !compatibleBallots(d.qBallot, d.pBallot)
		if !cond {
			log.Fatal("q ballot and p ballot invariant not satisfied")
		}
	}

	if d.hBallot != nil {
		if d.currentBallot == nil {
			log.Fatal("high ballot is not nil but current ballot is nil")
		}
		cond := compareBallots(d.hBallot, d.currentBallot) <= 0 && compatibleBallots(d.hBallot, d.currentBallot)
		if !cond {
			log.Fatal("current ballot and higher ballot invariant not satisfied")
		}
	}

	if d.cBallot != nil {
		if d.currentBallot == nil {
			log.Fatal("commit ballot is not nil but current ballot is nil")
		}
		cond := compareBallots(d.cBallot, d.hBallot) <= 0 && compatibleBallots(d.cBallot, d.hBallot)
		if !cond {
			log.Fatal("commit ballot and higher ballot invariant not satisfied")
		}
		cond = compareBallots(d.hBallot, d.currentBallot) <= 0 && compatibleBallots(d.hBallot, d.currentBallot)
		if !cond {
			log.Fatal("current ballot and higher ballot invariant not satisfied")
		}
	}

	switch d.currentPhase {
	case BallotPhasePrepare:
	case BallotPhaseConfirm:
		if d.cBallot == nil {
			log.Fatal("commit ballot should not be nil in confirm phase")
		}
	case BallotPhaseExternalize:
		if d.cBallot == nil {
			log.Fatal("commit ballot should not be nil in externalize phase")
		}
		if d.hBallot == nil {
			log.Fatal("higher ballot should not be nil in externalize phase")
		}
	default:
		log.Fatalf("invalid ballot phase: %d", d.currentPhase)
	}
}

// validate ballot by checking:
// 1. ballot counters are in expected states
// 2. ballot values are normal and satisfy consensus constraints
func (d *Decree) validateBallot(stmt *ultpb.Statement) error {
	if stmt == nil {
		return fmt.Errorf("ballot statement is nil")
	}

	fromSelf := d.nodeID == stmt.NodeID

	// value set for checking validity of ballot value
	values := mapset.NewSet()

	// TODO(bobonovski) check quorum sanity

	switch stmt.StatementType {
	case ultpb.StatementType_PREPARE:
		prepare, err := ultpb.DecodePrepare(stmt.Data)
		if err != nil {
			return fmt.Errorf("decode prepare statement failed: %v", err)
		}
		// checking counter sanity
		if !fromSelf && prepare.B.Counter == 0 {
			return errors.New("prepare ballot counter is zero and not self msg")
		}
		cond := compareBallots(prepare.Q, prepare.P) <= 0 && !compatibleBallots(prepare.Q, prepare.P)
		if prepare.Q != nil && prepare.P != nil && !cond {
			return errors.New("prepare q ballot and p ballot are not in expected states")
		}
		cond = prepare.HC == 0 || (prepare.P != nil && prepare.HC <= prepare.P.Counter)
		if !cond {
			return errors.New("prepare p ballot and higher counters are not in expected states")
		}
		cond = prepare.LC == 0 || (prepare.HC != 0 && prepare.HC <= prepare.B.Counter && prepare.LC <= prepare.HC)
		if !cond {
			return errors.New("prepare ballot counters are not in expected states")
		}
		// add value to set
		if prepare.B.Counter > 0 {
			values.Add(prepare.B.Value)
		}
		if prepare.P != nil {
			values.Add(prepare.P.Value)
		}
	case ultpb.StatementType_CONFIRM:
		confirm, err := ultpb.DecodeConfirm(stmt.Data)
		if err != nil {
			return fmt.Errorf("decode confirm statement failed: %v", err)
		}
		// check counter sanity
		if confirm.B.Counter == 0 {
			return fmt.Errorf("confirm current ballot counter should not be zero")
		}
		cond := confirm.LC <= confirm.HC && confirm.HC <= confirm.B.Counter
		if !cond {
			return fmt.Errorf("confirm ballot counters are not in expected states")
		}
		// add value to set
		values.Add(confirm.B.Value)
	case ultpb.StatementType_EXTERNALIZE:
		ext, err := ultpb.DecodeExternalize(stmt.Data)
		if err != nil {
			return fmt.Errorf("decode externalize statement failed: %v", err)
		}
		// check counter sanity
		cond := ext.B.Counter > 0 && ext.B.Counter <= ext.HC
		if !cond {
			return fmt.Errorf("externalize ballot counters are not in expected states")
		}
		// add value to set
		values.Add(ext.B.Value)
	default:
		log.Fatalf("invalid statement type: %v", stmt.StatementType)
	}

	// check values against consensus constraints
	var valueErr error
	for v := range values.Iter() {
		if err := d.validateConsensusValue(v.(string)); err != nil {
			valueErr = err
		}
	}

	return valueErr
}

// validate consensus value
func (d *Decree) validateConsensusValue(val string) error {
	vb, err := hex.DecodeString(val)
	if err != nil {
		return fmt.Errorf("decode hex string failed: %v", err)
	}

	_, err = ultpb.DecodeConsensusValue(vb)
	if err != nil {
		return fmt.Errorf("decode consensus value failed: %v", err)
	}

	// TODO(bobonovski) define maybe validate state

	return nil
}
