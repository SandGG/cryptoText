package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var text = "This is a hash"

	//checksum of the data
	var sum = sha256.Sum256([]byte(text))

	fmt.Println(text)
	fmt.Printf("%x\n", sum)
}
