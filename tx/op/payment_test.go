package op

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ultiledger/go-ultiledger/account"
	"github.com/ultiledger/go-ultiledger/db/memdb"
	"github.com/ultiledger/go-ultiledger/exchange"
	"github.com/ultiledger/go-ultiledger/ultpb"
)

func TestPaymentOp(t *testing.T) {
	memorydb := memdb.New()
	am := account.NewManager(memorydb, 100)
	em := exchange.NewManager(memorydb, am)

	// create source account
	err := am.CreateAccount(memorydb, srcAccount, 1000000, signer, 2)
	assert.Nil(t, err)

	// create destination account
	err = am.CreateAccount(memorydb, dstAccount, 10000, signer, 3)
	assert.Nil(t, err)

	memorytx, _ := memorydb.Begin()

	// create payment op
	paymentOp := Payment{
		AM:           am,
		EM:           em,
		SrcAccountID: srcAccount,
		DstAccountID: dstAccount,
		Asset:        &ultpb.Asset{AssetType: ultpb.AssetType_NATIVE, AssetName: "ULU", Issuer: signer},
		Amount:       int64(10000),
	}
	err = paymentOp.Apply(memorytx)
	assert.Nil(t, err)

	// check dst account
	dstAcc, err := am.GetAccount(memorytx, dstAccount)
	assert.Nil(t, err)
	assert.Equal(t, dstAcc.Balance, int64(20000))

	// check src account
	srcAcc, err := am.GetAccount(memorytx, srcAccount)
	assert.Nil(t, err)
	assert.Equal(t, srcAcc.Balance, int64(990000))
}
