package main

import "fmt"

func IntToHex(i int64) []byte {
	hex := fmt.Sprintf("%x", i)
	hbyte := []byte(hex)

	return hbyte
}
