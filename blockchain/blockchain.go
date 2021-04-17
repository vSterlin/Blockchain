package blockchain

import "github.com/vSterlin/blockchain/block"

type Blockchain struct {
	Chain []*block.Block
}

func NewBlockchain() (newBlockchain *Blockchain) {
	chain := []*block.Block{block.NewBlock(0, "Genesis", "")}
	newBlockchain = &Blockchain{
		Chain: chain,
	}
	return
}

func (bc *Blockchain) AddBlock(newBlock *block.Block) {
	bc.Chain = append(bc.Chain, newBlock)
}
