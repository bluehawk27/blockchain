package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// “target bits” is the block header storing the difficulty at which the block was mined
// TODO make this auto-adjusting
const targetBits = 24
const maxNonce = math.MaxInt64

// ProofOfWork : Proof of Work Object
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWOrk : initialize POW
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{
		Block:  b,
		Target: target,
	}

	return pow
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var (
		hashInt big.Int
		hash    [32]byte
	)
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.Block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

// Validate : Validates the POW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.Target) == -1
	return isValid
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data

}
