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

package tx

import (
	"fmt"
	"testing"

	"github.com/ultiledger/go-ultiledger/ultpb"

	"github.com/stretchr/testify/assert"
)

func TestTxHistory(t *testing.T) {
	txs := []*ultpb.Tx{
		&ultpb.Tx{AccountID: "A", SeqNum: uint64(1)},
		&ultpb.Tx{AccountID: "B", SeqNum: uint64(2)},
		&ultpb.Tx{AccountID: "C", SeqNum: uint64(3)},
		&ultpb.Tx{AccountID: "D", SeqNum: uint64(4)},
	}

	var keys []string

	txh := NewTxHistory()
	for i, tx := range txs {
		key := fmt.Sprintf("key-%d", i)
		txh.AddTx(key, tx)
		keys = append(keys, key)
	}

	assert.Equal(t, txh.Size(), 4)

	txh.DeleteTxList(keys)

	assert.Equal(t, txh.Size(), 0)
}
