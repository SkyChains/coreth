package miner

import (
	"math/big"

	"github.com/skychains/coreth/core/txpool"
	"github.com/skychains/coreth/core/types"
	"github.com/ethereum/go-ethereum/common"
)

type TransactionsByPriceAndNonce = transactionsByPriceAndNonce

func NewTransactionsByPriceAndNonce(signer types.Signer, txs map[common.Address][]*txpool.LazyTransaction, baseFee *big.Int) *TransactionsByPriceAndNonce {
	return newTransactionsByPriceAndNonce(signer, txs, baseFee)
}
