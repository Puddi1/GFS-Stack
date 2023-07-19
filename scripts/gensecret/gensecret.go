package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main () {
	buf := make([]byte, 32)
	rand.Read(buf)
	fmt.Println(hex.EncodeToString(buf)[:32])
}