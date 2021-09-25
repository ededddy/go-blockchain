package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Invalid Address")
	}
	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Done!")
}
