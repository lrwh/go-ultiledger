package consensus

import (
	"errors"
	"fmt"
	"math"

	"github.com/deckarep/golang-set"

	"github.com/ultiledger/go-ultiledger/ultpb"
)

// Decree is an abstractive decision the consensus engine
// should reach in each round
type Decree struct {
	index uint64

	// nodeID of this node
	nodeID string

	// nomination round
	round int

	// latest nomination
	latestNomination *ultpb.Nomination

	votes       mapset.Set
	accepts     mapset.Set
	candidates  mapset.Set
	nominations map[string]*ultpb.Nomination

	// channel for sending statements
	statementChan chan *ultpb.Statement
}

func NewDecree(idx uint64, nodeID string, stmtC chan *ultpb.Statement) *Decree {
	d := &Decree{
		index:         idx,
		nodeID:        nodeID,
		round:         0,
		votes:         mapset.NewSet(),
		accepts:       mapset.NewSet(),
		candidates:    mapset.NewSet(),
		nominations:   make(map[string]*ultpb.Nomination),
		statementChan: stmtC,
	}
	return d
}

// nominate a consensus value for this slot
func (d *Decree) Nominate(quorum *ultpb.Quorum, quorumHash, prevHash, currHash string) error {
	d.round++
	// TODO(bobonovski) compute leader weights
	d.votes.Add(currHash) // For test

	if err := d.sendNomination(quorum, quorumHash); err != nil {
		return fmt.Errorf("send nomination failed: %v", err)
	}
	return nil
}

// receive nomination from peers or local node
func (d *Decree) recvNomination(nodeID string, quorum *ultpb.Quorum, quorumHash string, nom *ultpb.Nomination) error {
	d.addNomination(d.nodeID, nom)
	acceptUpdated, candidateUpdated, err := d.promoteVotes(quorum, nom)
	if err != nil {
		return fmt.Errorf("promote votes failed: %v", err)
	}

	// send new nomination if votes changed
	if acceptUpdated {
		d.sendNomination(quorum, quorumHash)
	}

	// start balloting if candidates changed
	if candidateUpdated {
		// TODO(bobonovski) balloting
	}

	return nil
}

// assemble a nomination and broadcast it to other peers
func (d *Decree) sendNomination(quorum *ultpb.Quorum, quorumHash string) error {
	// create an abstract nomination statement
	nom := &ultpb.Nomination{
		QuorumHash: quorumHash,
	}
	for vote := range d.votes.Iter() {
		nom.VoteList = append(nom.VoteList, vote.(string))
	}
	for accept := range d.accepts.Iter() {
		nom.AcceptList = append(nom.AcceptList, accept.(string))
	}

	if err := d.recvNomination(d.nodeID, quorum, quorumHash, nom); err != nil {
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
			SlotIndex:     d.index,
			Data:          nomBytes,
		}
		d.statementChan <- stmt
	}

	return nil
}

// check whether the input nomination is valid and newer
func (d *Decree) addNomination(nodeID string, newNom *ultpb.Nomination) error {
	// check validity of votes and accepts
	if len(newNom.VoteList)+len(newNom.AcceptList) == 0 {
		return errors.New("vote and accept list is empty")
	}

	// check whether the existing nomination of the remote node
	// is the proper subset of the new nomination
	if nom, ok := d.nominations[nodeID]; ok {
		if isNewerNomination(nom, newNom) {
			d.nominations[nodeID] = newNom
		}
	}

	return nil
}

func isNewerNomination(anom *ultpb.Nomination, bnom *ultpb.Nomination) bool {
	if anom == nil && bnom != nil {
		return true
	}

	if !IsProperSubset(anom.VoteList, bnom.VoteList) {
		// TODO(bobonovski) more elaborate check like interset?
		return false
	}

	if !IsProperSubset(anom.AcceptList, bnom.AcceptList) {
		return false
	}

	return true
}

// check whether the input node set form V-blocking for input quorum
func isVblocking(quorum *ultpb.Quorum, nodeSet mapset.Set) bool {
	qsize := float64(len(quorum.Validators) + len(quorum.NestQuorums))
	threshold := int(math.Ceil(qsize * (1.0 - quorum.Threshold)))

	for _, vid := range quorum.Validators {
		if nodeSet.Contains(vid) {
			threshold = threshold - 1
		}
		if threshold == 0 {
			return true
		}
	}

	for _, nq := range quorum.NestQuorums {
		if isVblocking(nq, nodeSet) {
			threshold = threshold - 1
		}
		if threshold == 0 {
			return true
		}
	}

	return false
}

// check whether the input node set form quorum slice for input quorum
func isQuorumSlice(quorum *ultpb.Quorum, nodeSet mapset.Set) bool {
	qsize := float64(len(quorum.Validators) + len(quorum.NestQuorums))
	threshold := int(math.Ceil(qsize * quorum.Threshold))

	for _, vid := range quorum.Validators {
		if nodeSet.Contains(vid) {
			threshold = threshold - 1
		}
		if threshold == 0 {
			return true
		}
	}

	for _, nq := range quorum.NestQuorums {
		if isVblocking(nq, nodeSet) {
			threshold = threshold - 1
		}
		if threshold == 0 {
			return true
		}
	}

	return false
}

// find set of nodes claimed to accept the value
func findAcceptNodes(v string, noms map[string]*ultpb.Nomination) mapset.Set {
	nodeSet := mapset.NewSet()

	for k, nom := range noms {
		for _, av := range nom.AcceptList {
			if v == av {
				nodeSet.Add(k)
				break
			}
		}
	}

	return nodeSet
}

// find set of nodes claimed to vote or accept the value
func findVoteOrAcceptNodes(v string, noms map[string]*ultpb.Nomination) mapset.Set {
	nodeSet := mapset.NewSet()

	for k, nom := range noms {
		for _, vv := range nom.VoteList {
			if v == vv {
				nodeSet.Add(k)
				break
			}
		}
		for _, av := range nom.AcceptList {
			if v == av {
				nodeSet.Add(k)
				break
			}
		}
	}

	return nodeSet
}

// try to promote votes to accepts by checking two conditions:
//   1. whether the votes form V-blocking
//   2. whether all the nodes in the quorum have voted
// then try to promote accepts to candidates by checking:
//   1. whether all the nodes in the quorum have accepted
func (d *Decree) promoteVotes(quorum *ultpb.Quorum, newNom *ultpb.Nomination) (bool, bool, error) {
	acceptUpdated := false
	for _, vote := range newNom.VoteList {
		if d.accepts.Contains(vote) {
			continue
		}

		// use federated vote to promote value
		ns := findAcceptNodes(vote, d.nominations)
		if !isVblocking(quorum, ns) {
			nset := findVoteOrAcceptNodes(vote, d.nominations)
			if !isQuorumSlice(quorum, nset) { // TODO(bobonovski) trim nset to contain only other quorums
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
		if isQuorumSlice(quorum, ns) {
			d.candidates.Add(accept)
			candidateUpdated = true
		}
	}

	return acceptUpdated, candidateUpdated, nil
}