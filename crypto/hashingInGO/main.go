package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("Hello, Hezron!")
	hash := sha256.Sum256(data)
	fmt.Printf("SHA-256: %x\n", hash)
}
