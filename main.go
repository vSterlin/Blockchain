package main

import (
	"fmt"
	"time"

	"github.com/vSterlin/blockchain/block"
	"github.com/vSterlin/blockchain/blockchain"
)

func main() {
	blockchain := blockchain.NewBlockchain()

	blockchain.AddBlock(&block.Block{
		Index:        uint(len(blockchain.Chain)),
		Timestamp:    time.Now(),
		Data:         "abc",
		PreviousHash: "abcbcb",
	})
	for _, block := range blockchain.Chain {
		fmt.Printf("%+v\n", block)

	}
}
