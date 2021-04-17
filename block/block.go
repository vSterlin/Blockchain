package block

import (
	"time"
)

type Block struct {
	Index     uint
	Timestamp time.Time
	// data         data
	Data         string
	PreviousHash string
	Hash         string
}

func NewBlock(index uint, data string, previousHash string) (newBlock *Block) {
	newBlock = &Block{
		Index:        index,
		Timestamp:    time.Now(),
		Data:         data,
		PreviousHash: previousHash,
	}
	return
}

// type data struct {
// 	text string
// }
