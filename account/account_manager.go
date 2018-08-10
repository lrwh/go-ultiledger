package account

import (
	"go.uber.org/zap"

	"github.com/ultiledger/go-ultiledger/crypto"
	pb "github.com/ultiledger/go-ultiledger/ultpb"
)

// AccountManager manages the creation of accounts
type accountManager struct {
	logger *zap.SugaredLogger
	// master account
	master *pb.Account
}

func NewAccountManager(l *zap.SugaredLogger) *accountManager {
	return &accountManager{logger: l}
}

// Create master account with native asset (ULT) and initial balances
func (am *accountManager) CreateMasterAccount(networkID []byte, balance uint64) error {
	pubKey, privKey, err := crypto.GenerateKeypairFromSeed(networkID)
	if err != nil {
		return err
	}
	am.logger.Infof("master private key (seed) is %s", privKey)

	ult := &pb.Asset{
		AssetType: pb.Asset_NATIVE,
		AssetName: "ULT",
		Signer:    pubKey,
		Balance:   balance,
	}
	acc := &pb.Account{
		AccountID: pubKey,
		Assets:    []*pb.Asset{ult},
		Signer:    pubKey,
	}
	am.master = acc

	return nil
}
