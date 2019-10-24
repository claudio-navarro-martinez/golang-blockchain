package main

import (
	"fmt"
	"rsc.io/quote"
	"github.com/claudio-navarro-martinez/golang-blockchain/blockchain"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(blockchain.Kkmodulo)
	blockchain.Imprimekkmodulo(10)
}