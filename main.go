package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("send 1 BTC to Scott")
	bc.AddBlock("send 2 BTC to Ben")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
