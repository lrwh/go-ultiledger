// Copyright 2019 The go-ultiledger Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consensus

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/deckarep/golang-set"
	lru "github.com/hashicorp/golang-lru"
	b58 "github.com/mr-tron/base58/base58"

	"github.com/ultiledger/go-ultiledger/db"
	"github.com/ultiledger/go-ultiledger/ledger"
	"github.com/ultiledger/go-ultiledger/log"
	"github.com/ultiledger/go-ultiledger/ultpb"
)

// ValidatorContext contains contextual information validator needs.
type ValidatorContext struct {
	Database           db.Database // database instance
	LM                 *ledger.Manager
	MaxDecrees         uint64
	QuorumDownloadChan chan<- string
	TxSetDownloadChan  chan<- string
}

func ValidateValidatorContext(vc *ValidatorContext) error {
	if vc == nil {
		return errors.New("validator context is nil")
	}
	if vc.Database == nil {
		return errors.New("db instance is nil")
	}
	if vc.LM == nil {
		return errors.New("ledger manager is nil")
	}
	if vc.MaxDecrees == 0 {
		return errors.New("max decrees is zero")
	}
	if vc.QuorumDownloadChan == nil {
		return errors.New("quorum download chan is nil")
	}
	if vc.TxSetDownloadChan == nil {
		return errors.New("txset download chan is nil")
	}
	return nil
}

// Validator validates incoming consensus messages and
// requests missing information of transaction and quorum
// from other peers.
type Validator struct {
	database db.Database
	bucket   string

	lm *ledger.Manager

	// Maximum decrees to cache.
	maxDecrees uint64

	// Statements which are downloading.
	rwm       sync.RWMutex
	downloads map[string]*Statement

	// Statements of each decree index.
	decreeStmts map[uint64][]*Statement

	// Max sequence number reiceived.
	maxSeqNum uint64

	// Set of received statements.
	statements mapset.Set

	// Cache of quorum with quorum hash as the key.
	quorumCache *lru.Cache

	// Channel for sending quorum download task.
	quorumDownloadChan chan<- string
	// Channel for sending txset download task.
	txsetDownloadChan chan<- string

	stopChan chan struct{}

	// Channel for notifying consensus engine that the statement is ready.
	readyChan chan *Statement
	// Channel for dispatching statements with the next ledger sequence.
	dispatchChan chan *Statement
	// Channel for downloding missing information of the statement.
	downloadChan chan *Statement
}

func NewValidator(ctx *ValidatorContext) *Validator {
	if err := ValidateValidatorContext(ctx); err != nil {
		log.Fatalf("validator context is invalid: %v", err)
	}

	v := &Validator{
		database:           ctx.Database,
		bucket:             "VALIDATOR",
		lm:                 ctx.LM,
		maxDecrees:         ctx.MaxDecrees,
		downloads:          make(map[string]*Statement),
		decreeStmts:        make(map[uint64][]*Statement),
		maxSeqNum:          uint64(0),
		statements:         mapset.NewSet(),
		quorumDownloadChan: ctx.QuorumDownloadChan,
		txsetDownloadChan:  ctx.TxSetDownloadChan,
		stopChan:           make(chan struct{}),
		readyChan:          make(chan *Statement, 100),
		downloadChan:       make(chan *Statement, 100),
		dispatchChan:       make(chan *Statement, 100),
	}

	err := v.database.NewBucket(v.bucket)
	if err != nil {
		log.Fatalf("create validator bucket failed: %v", err)
	}

	qc, err := lru.New(1000)
	if err != nil {
		log.Fatalf("create quorum LRU cache failed: %v", err)
	}
	v.quorumCache = qc

	// Listening for download tasks.
	go v.download()
	// Monitor for downloaded statements.
	go v.monitor()
	// Dispatch ready statments.
	go v.dispatch()

	return v
}

// Stop the validator.
func (v *Validator) Stop() {
	close(v.stopChan)
}

// Ready returns ready statements with full information.
func (v *Validator) Ready() <-chan *Statement {
	return v.readyChan
}

// Recv checks whether the input statement has the complete
// information including quorum and txset for the consensus
// engine to process. If all the necessary information are
// present, it will directly return the statement to ready
// channel. Otherwise, it will try to download the missing
// part of information from peers.
func (v *Validator) Recv(stmt *Statement) error {
	if stmt == nil {
		return nil
	}

	// Filter the duplicate statement using the statement hash.
	hash, err := ultpb.SHA256Hash(stmt)
	if err != nil {
		return fmt.Errorf("compute statement hash failed: %v", err)
	}
	if v.statements.Contains(hash) {
		return nil
	}
	v.statements.Add(hash)

	// Check whether the statement has all the information.
	valid, err := v.validate(stmt)
	if err != nil {
		return fmt.Errorf("validate statement failed: %v", err)
	}

	if valid {
		v.dispatchChan <- stmt
	} else {
		v.downloadChan <- stmt
		// Save the ongoing downloading statement.
		v.rwm.Lock()
		v.downloads[hash] = stmt
		v.rwm.Unlock()
	}

	return nil
}

// RecvQuorum receives downloaded quorum and save it.
func (v *Validator) RecvQuorum(quorumHash string, quorum *Quorum) error {
	if err := ValidateQuorum(quorum, 0, false); err != nil {
		return err
	}

	// Encode the quorum to pb format.
	qb, err := ultpb.Encode(quorum)
	if err != nil {
		return fmt.Errorf("encode quorum failed: %v", err)
	}

	err = v.database.Put(v.bucket, []byte(quorumHash), qb)
	if err != nil {
		return fmt.Errorf("save quorum to db failed: %v", err)
	}

	v.quorumCache.Add(quorumHash, quorum)

	return nil
}

// RecvTxSet receives downloaded txset and save it.
func (v *Validator) RecvTxSet(txsetHash string, txset *TxSet) error {
	err := v.lm.AddTxSet(txsetHash, txset)
	if err != nil {
		return err
	}
	return nil
}

// Get the txset from the ledger manager.
func (v *Validator) GetTxSet(txsetHash string) (*TxSet, error) {
	txset, err := v.lm.GetTxSet(txsetHash)
	if err != nil {
		return nil, err
	}
	return txset, nil
}

// Get the quorum of the corresponding quorum hash,
func (v *Validator) GetQuorum(quorumHash string) (*Quorum, error) {
	if q, ok := v.quorumCache.Get(quorumHash); ok {
		return q.(*Quorum), nil
	}

	qb, err := v.database.Get(v.bucket, []byte(quorumHash))
	if err != nil {
		return nil, fmt.Errorf("get quorum from db failed: %v", err)
	}
	if qb == nil {
		return nil, nil
	}

	quorum, err := ultpb.DecodeQuorum(qb)
	if err != nil {
		return nil, fmt.Errorf("decode quorum failed: %v", err)
	}

	// Cache the quorum.
	v.quorumCache.Add(quorumHash, quorum)

	return quorum, nil
}

// Monitor downloaded statements and send them to dispatch channel.
func (v *Validator) monitor() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			v.rwm.RLock()
			for h, stmt := range v.downloads {
				valid, _ := v.validate(stmt)
				if valid {
					v.dispatchChan <- stmt
					delete(v.downloads, h)
				}
			}
			v.rwm.RUnlock()
		case <-v.stopChan:
			return
		}
	}
}

// Dispatch the statements with the next ledger sequence to the ready channel.
func (v *Validator) dispatch() {
	var nextSeqNum uint64
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			nextSeqNum = v.lm.NextLedgerHeaderSeq()
			// Future statements are tried to dispatched in case that
			// the local node did not receive necessary statements with
			// the next sequence number.
			for i := nextSeqNum; i <= v.maxSeqNum; i++ {
				if len(v.decreeStmts[i]) == 0 {
					continue
				}
				// First dispatch cached statments.
				for _, stmt := range v.decreeStmts[i] {
					v.readyChan <- stmt
				}
				// Clear dispatched statements
				v.decreeStmts[i] = v.decreeStmts[i][:0]
			}
			// Remove old statements.
			for {
				if nextSeqNum <= v.maxDecrees {
					break
				}
				index := nextSeqNum - v.maxDecrees
				if _, ok := v.decreeStmts[index]; ok {
					delete(v.decreeStmts, index)
					index -= 1
				} else {
					break
				}
			}
		case stmt := <-v.dispatchChan:
			nextSeqNum = v.lm.NextLedgerHeaderSeq()
			if stmt.Index == nextSeqNum {
				v.readyChan <- stmt
			} else {
				v.decreeStmts[stmt.Index] = append(v.decreeStmts[stmt.Index], stmt)
				if stmt.Index > v.maxSeqNum {
					v.maxSeqNum = stmt.Index
				}
			}
		case <-v.stopChan:
			return
		}
	}
}

// Download missing info of the statement.
func (v *Validator) download() {
	for {
		select {
		case stmt := <-v.downloadChan:
			// No need to check error as we have already passed the validity checks.
			quorumHash, _ := extractQuorumHash(stmt)
			quorum, _ := v.GetQuorum(quorumHash)
			if quorum == nil {
				v.quorumDownloadChan <- quorumHash
			}

			txsetHashes, _ := extractTxSetHash(stmt)
			for _, txsetHash := range txsetHashes {
				txset, _ := v.lm.GetTxSet(txsetHash)
				if txset == nil {
					v.txsetDownloadChan <- txsetHash
				}
			}
		case <-v.stopChan:
			return
		}
	}
}

// Validate statement by checking whether we have its quorum and txset.
func (v *Validator) validate(stmt *Statement) (bool, error) {
	quorumHash, err := extractQuorumHash(stmt)
	if err != nil {
		return false, fmt.Errorf("extract quorum hash from statement failed: %v", err)
	}

	quorum, err := v.GetQuorum(quorumHash)
	if err != nil {
		return false, fmt.Errorf("query quorum failed: %v", err)
	}
	if quorum == nil {
		return false, nil
	}

	txsetHashes, err := extractTxSetHash(stmt)
	if err != nil {
		return false, fmt.Errorf("extract tx list hash from statement failed: %v", err)
	}

	for _, txsetHash := range txsetHashes {
		txset, err := v.lm.GetTxSet(txsetHash)
		if err != nil {
			return false, fmt.Errorf("query txset failed: %v", err)
		}
		if txset == nil {
			return false, nil
		}
	}

	return true, nil
}

// Extract quorum hash from statement.
func extractQuorumHash(stmt *Statement) (string, error) {
	if stmt == nil {
		return "", errors.New("statement is nil")
	}
	var hash string
	switch stmt.StatementType {
	case ultpb.StatementType_NOMINATE:
		nom := stmt.GetNominate()
		hash = nom.QuorumHash
	case ultpb.StatementType_PREPARE:
		prepare := stmt.GetPrepare()
		hash = prepare.QuorumHash
	case ultpb.StatementType_CONFIRM:
		confirm := stmt.GetConfirm()
		hash = confirm.QuorumHash
	case ultpb.StatementType_EXTERNALIZE:
		ext := stmt.GetExternalize()
		hash = ext.QuorumHash
	default:
		log.Fatal(ErrUnknownStmtType)
	}
	return hash, nil
}

// Extract list of tx set hash from statement.
func extractTxSetHash(stmt *Statement) ([]string, error) {
	if stmt == nil {
		return nil, errors.New("statement is nil")
	}
	var hashes []string
	switch stmt.StatementType {
	case ultpb.StatementType_NOMINATE:
		nom := stmt.GetNominate()
		for _, v := range nom.VoteList {
			b, err := b58.Decode(v)
			if err != nil {
				return nil, fmt.Errorf("decode hex string failed: %v", err)
			}
			cv, err := ultpb.DecodeConsensusValue(b)
			if err != nil {
				return nil, fmt.Errorf("decode consensus value failed: %v", err)
			}
			hashes = append(hashes, cv.TxSetHash)
		}
		for _, a := range nom.AcceptList {
			b, err := b58.Decode(a)
			if err != nil {
				return nil, fmt.Errorf("decode hex string failed: %v", err)
			}
			cv, err := ultpb.DecodeConsensusValue(b)
			if err != nil {
				return nil, fmt.Errorf("decode consensus value failed: %v", err)
			}
			hashes = append(hashes, cv.TxSetHash)
		}
	case ultpb.StatementType_PREPARE:
		fallthrough
	case ultpb.StatementType_CONFIRM:
		fallthrough
	case ultpb.StatementType_EXTERNALIZE:
		ballot := getWorkingBallot(stmt)
		b, err := b58.Decode(ballot.Value)
		if err != nil {
			return nil, fmt.Errorf("decode hex string failed: %v", err)
		}
		cv, err := ultpb.DecodeConsensusValue(b)
		if err != nil {
			return nil, fmt.Errorf("decode consensus value failed: %v", err)
		}
		hashes = append(hashes, cv.TxSetHash)
	default:
		return nil, errors.New("unknown statement type")
	}
	return hashes, nil
}
