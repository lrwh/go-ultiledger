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

package op

import (
	"errors"
	"fmt"
	"math"

	pb "github.com/golang/protobuf/proto"

	"github.com/ultiledger/go-ultiledger/account"
	"github.com/ultiledger/go-ultiledger/db"
	"github.com/ultiledger/go-ultiledger/exchange"
	"github.com/ultiledger/go-ultiledger/ultpb"
)

var (
	ErrInvalidPaymentAmount = errors.New("invalid payment amount")
	ErrInvalidAccountID     = errors.New("invalid accountID")
	ErrPaymentNotAuthorized = errors.New("payment is not authorized")
)

func validateAsset(asset *ultpb.Asset) error {
	if asset == nil {
		return errors.New("asset is nil")
	}
	if asset.AssetType == ultpb.AssetType_NATIVE {
		return nil
	}
	if len(asset.AssetName) <= 0 || len(asset.AssetName) > 4 {
		return errors.New("invalid asset name")
	}
	return nil
}

// Peer to peer payment in specified asset.
type Payment struct {
	AM           *account.Manager
	EM           *exchange.Manager
	SrcAccountID string
	DstAccountID string
	Asset        *ultpb.Asset
	Amount       int64
}

func (p *Payment) Apply(dt db.Tx) error {
	if err := validateAsset(p.Asset); err != nil {
		return fmt.Errorf("validate payment asset failed: %v", err)
	}
	if p.Amount == 0 {
		return ErrInvalidPaymentAmount
	}
	if p.SrcAccountID == "" || p.DstAccountID == "" {
		return ErrInvalidAccountID
	}

	// Payment is equivalent to path payment with one hop,
	// so we construct a path payment to reuse the logic.
	pp := &PathPayment{
		AM:           p.AM,
		EM:           p.EM,
		SrcAccountID: p.SrcAccountID,
		SrcAsset:     p.Asset,
		SrcAmount:    p.Amount,
		DstAccountID: p.DstAccountID,
		DstAsset:     p.Asset,
		DstAmount:    p.Amount,
	}
	if err := pp.Apply(dt); err != nil {
		return err
	}

	return nil
}

// Path payment from source asset to destination asset.
type PathPayment struct {
	AM           *account.Manager
	EM           *exchange.Manager
	SrcAccountID string
	SrcAsset     *ultpb.Asset
	SrcAmount    int64
	DstAccountID string
	DstAsset     *ultpb.Asset
	DstAmount    int64
	Path         []*ultpb.Asset
}

func (pp *PathPayment) Apply(dt db.Tx) error {
	if err := validateAsset(pp.SrcAsset); err != nil {
		return fmt.Errorf("validate src payment asset failed: %v", err)
	}
	if err := validateAsset(pp.DstAsset); err != nil {
		return fmt.Errorf("validate dst payment asset failed: %v", err)
	}
	for _, a := range pp.Path {
		if err := validateAsset(a); err != nil {
			return fmt.Errorf("validate path payment asset failed: %v", err)
		}
	}
	if pp.SrcAccountID == "" || pp.DstAccountID == "" {
		return ErrInvalidAccountID
	}
	if pp.SrcAmount == 0 || pp.DstAmount == 0 {
		return ErrInvalidPaymentAmount
	}

	// Save the last asset and amount exchanged.
	asset, amount := pp.DstAsset, pp.DstAmount

	// Build asset exchange path.
	var path []*ultpb.Asset
	path = append(path, pp.SrcAsset)
	path = append(path, pp.Path...)

	dstAccount, err := pp.AM.GetAccount(dt, pp.DstAccountID)
	if err != nil {
		return fmt.Errorf("get dst account failed: %v", err)
	}
	if dstAccount == nil {
		return ErrAccountNotExist
	}

	if asset.AssetType == ultpb.AssetType_NATIVE {
		if err := pp.AM.UpdateBalance(dstAccount, amount); err != nil {
			return fmt.Errorf("add balance failed: %v", err)
		}
		if err := pp.AM.SaveAccount(dt, dstAccount); err != nil {
			return fmt.Errorf("save account failed: %v", err)
		}
	} else {
		issuer, err := pp.AM.GetAccount(dt, asset.Issuer)
		if err != nil {
			return fmt.Errorf("get asset issuer failed: %v", err)
		}
		if issuer == nil {
			return ErrAccountNotExist
		}

		trust, err := pp.AM.GetTrust(dt, pp.DstAccountID, asset)
		if err != nil {
			return fmt.Errorf("get dst trust failed: %v", err)
		}

		if trust.Authorized == 0 {
			return ErrPaymentNotAuthorized
		}

		if err := pp.AM.UpdateTrustBalance(trust, pp.DstAmount); err != nil {
			return fmt.Errorf("add trust balance failed: %v", err)
		}

		if err := pp.AM.SaveTrust(dt, trust); err != nil {
			return fmt.Errorf("save trust failed: %v", err)
		}
	}

	//TODO(bobonovski) exchange assets in backward order
	for i := len(path) - 1; i >= 0; i-- {
		if pb.Equal(path[i], asset) {
			continue
		}
		// Check whether asset issuer exists.
		if path[i].AssetType != ultpb.AssetType_NATIVE {
			issuer, err := pp.AM.GetAccount(dt, path[i].Issuer)
			if err != nil {
				return fmt.Errorf("get issuer account failed: %v", err)
			}
			if issuer == nil {
				return ErrAccountNotExist
			}
		}
		// Exchange assets.
		order := &exchange.Order{
			SellAsset:    path[i],
			MaxSellAsset: math.MaxInt64,
			BuyAsset:     asset,
			MaxBuyAsset:  amount,
		}
		err = pp.EM.FillOrder(dt, order)
		if err != nil {
			return fmt.Errorf("exchange assets failed: %v", err)
		}
		asset = path[i]
		amount = order.SellAssetSold
	}

	if amount > pp.SrcAmount {
		return errors.New("deduced src payment amount is over the limit")
	}

	// Update source account balance.
	if asset.AssetType == ultpb.AssetType_NATIVE {
		srcAccount, err := pp.AM.GetAccount(dt, pp.SrcAccountID)
		if err != nil {
			return fmt.Errorf("load source account failed: %v", err)
		}
		if srcAccount == nil {
			return ErrAccountNotExist
		}
		if err := pp.AM.UpdateBalance(srcAccount, -amount); err != nil {
			return err
		}
		if err := pp.AM.SaveAccount(dt, srcAccount); err != nil {
			return err
		}
	} else {
		issuer, err := pp.AM.GetAccount(dt, asset.Issuer)
		if err != nil {
			return fmt.Errorf("get asset issuer failed: %v", err)
		}
		if issuer == nil {
			return ErrAccountNotExist
		}
		trust, err := pp.AM.GetTrust(dt, pp.SrcAccountID, asset)
		if err != nil {
			return fmt.Errorf("get dst trust failed: %v", err)
		}
		if trust.Authorized == 0 {
			return ErrPaymentNotAuthorized
		}
		if err := pp.AM.UpdateTrustBalance(trust, -amount); err != nil {
			return fmt.Errorf("add trust balance failed: %v", err)
		}
		if err := pp.AM.SaveTrust(dt, trust); err != nil {
			return fmt.Errorf("save trust failed: %v", err)
		}
	}

	return nil
}
