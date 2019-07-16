package op

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ultiledger/go-ultiledger/account"
	"github.com/ultiledger/go-ultiledger/db/memdb"
)

const (
	signer     = "SIGNER"
	srcAccount = "srcAccount"
	dstAccount = "dstAccount"
)

func TestAccountOp(t *testing.T) {
	memorydb := memdb.New()
	am := account.NewManager(memorydb, 100)

	// create source account
	err := am.CreateAccount(memorydb, srcAccount, 1000000, signer, 2)
	assert.Nil(t, err)

	memorytx, _ := memorydb.Begin()

	// create account op
	accountOp := CreateAccount{
		AM:           am,
		SrcAccountID: srcAccount,
		DstAccountID: dstAccount,
		Balance:      100000,
		BaseReserve:  100,
		SeqNum:       3,
	}
	err = accountOp.Apply(memorytx)
	assert.Nil(t, err)

	// check dst account
	dstAcc, err := am.GetAccount(memorytx, dstAccount)
	assert.Nil(t, err)
	assert.Equal(t, dstAcc.Balance, int64(100000))

	// check src account
	srcAcc, err := am.GetAccount(memorytx, srcAccount)
	assert.Nil(t, err)
	assert.Equal(t, srcAcc.Balance, int64(900000))
}
