// (c) 2021-2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package handlers

import (
	"github.com/skychains/coreth/core/state/snapshot"
	"github.com/skychains/coreth/core/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ BlockProvider    = &TestBlockProvider{}
	_ SnapshotProvider = &TestSnapshotProvider{}
)

type TestBlockProvider struct {
	GetBlockFn func(common.Hash, uint64) *types.Block
}

func (t *TestBlockProvider) GetBlock(hash common.Hash, number uint64) *types.Block {
	return t.GetBlockFn(hash, number)
}

type TestSnapshotProvider struct {
	Snapshot *snapshot.Tree
}

func (t *TestSnapshotProvider) Snapshots() *snapshot.Tree {
	return t.Snapshot
}
