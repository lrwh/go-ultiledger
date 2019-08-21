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

package util

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringArgSort(t *testing.T) {
	names := []string{"Bob", "Alex", "Zara", "July"}
	s := NewStringSlice(true, names...)
	sort.Sort(s)

	assert.Equal(t, s.Idx, []int{2, 3, 0, 1})
}

func TestIntArgSort(t *testing.T) {
	ints := []int{2, 4, 6, 3, 2, 1}
	is := NewIntSlice(false, ints...)
	sort.Sort(is)

	assert.Equal(t, is.Idx, []int{5, 0, 4, 3, 1, 2})
}

func TestFloat64ArgSort(t *testing.T) {
	floats := []float64{2.0, 4.0, 7.0, 1.0, 2.5}
	fs := NewFloat64Slice(true, floats...)
	sort.Sort(fs)

	assert.Equal(t, fs.Idx, []int{2, 1, 4, 0, 3})
}
