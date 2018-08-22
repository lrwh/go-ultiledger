package consensus

import (
	"errors"

	"github.com/deckarep/golang-set"
	"go.uber.org/zap"

	"github.com/ultiledger/go-ultiledger/account"
	"github.com/ultiledger/go-ultiledger/crypto"
	"github.com/ultiledger/go-ultiledger/db"
	"github.com/ultiledger/go-ultiledger/ledger"
	"github.com/ultiledger/go-ultiledger/peer"
	"github.com/ultiledger/go-ultiledger/types"
	pb "github.com/ultiledger/go-ultiledger/ultpb"
)

var (
	ErrInvalidTx           = errors.New("invalid transaction")
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidSeqNum       = errors.New("invalid sequence number")
)

// Engine is responsible for coordinating between upstream
// events and underlying consensus protocol
type Engine struct {
	store  db.DB
	bucket string

	logger *zap.SugaredLogger

	pm *peer.Manager
	am *account.Manager

	// consensus quorum
	quorum *pb.Quorum

	// latest voted slot
	latestSlotIdx uint64

	// vote round
	roundNum uint32

	// consensus protocol
	cp *ucp

	// transactions waiting to be include in the ledger
	txSet mapset.Set

	// accountID to txList map
	txMap map[string]*types.TxHistory

	nominateChan chan string
}

func NewEngine(d db.DB, l *zap.SugaredLogger, pm *peer.Manager, am *account.Manager) *Engine {
	e := &Engine{
		store:         d,
		bucket:        "ENGINE",
		logger:        l,
		pm:            pm,
		am:            am,
		latestSlotIdx: uint64(0),
		cp:            newUCP(d, l),
		txSet:         mapset.NewSet(),
		txMap:         make(map[string]*types.TxHistory),
	}
	err := e.store.CreateBucket(e.bucket)
	if err != nil {
		e.logger.Fatal(err)
	}
	return e
}

func (e *Engine) Start(stopChan chan struct{}) {
	go func() {
		for {
			select {
			case <-stopChan:
				return
			}
		}
	}()
}

// Add transaction to internal pending set
func (e *Engine) AddTx(tx *pb.Tx) error {
	h, err := crypto.SHA256HashPb(tx)
	if err != nil {
		return err
	}
	if e.txSet.Contains(h) {
		return errors.New("duplicate transaction")
	}
	// get the account information
	acc, err := e.am.GetAccount(tx.AccountID)
	if err != nil {
		e.logger.Warnw("cannot find the account", "AccountID", tx.AccountID)
		return err
	}
	// compute the total fees and max sequence number
	totalFees := tx.Fee
	maxSeq := tx.SequenceNumber
	if history, ok := e.txMap[tx.AccountID]; ok {
		totalFees += history.TotalFees
		maxSeq = MaxUint64(maxSeq, history.MaxSeqNum)
	} else {
		e.txMap[tx.AccountID] = types.NewTxHistory()
	}
	// check whether tx sequence number is larger than existing one
	if maxSeq > tx.SequenceNumber {
		return ErrInvalidSeqNum
	}
	// check whether the accounts has sufficient balance
	if acc.Balance-ledger.GenesisBaseReserve*uint64(acc.ItemCount) < totalFees {
		return ErrInsufficientBalance
	}
	e.txMap[tx.AccountID].AddTx(tx)
	e.txSet.Add(h)
	return nil
}
