package main

// BlockChain : BlockChain object
type BlockChain struct {
	Blocks []*Block
}

// NewBlockchain : instantiate a new BlockChain
func NewBlockchain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}

// AddBlock : Add new block to the chain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
