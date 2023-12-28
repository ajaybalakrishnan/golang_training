package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to Math in golang")

	// rand.NewSource(time.Now().UnixNano())
	// fmt.Println("Random Number", 1+rand.Intn(5))

	num, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Random Number from Crypto package", num)
}
