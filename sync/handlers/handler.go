// (c) 2021-2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package handlers

import (
	"github.com/skychains/coreth/core/state/snapshot"
	"github.com/skychains/coreth/core/types"
	"github.com/ethereum/go-ethereum/common"
)

type BlockProvider interface {
	GetBlock(common.Hash, uint64) *types.Block
}

type SnapshotProvider interface {
	Snapshots() *snapshot.Tree
}

type SyncDataProvider interface {
	BlockProvider
	SnapshotProvider
}
