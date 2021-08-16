package main

import (
	"fmt"

	"github.com/whguddnr9/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "nico", Balance: 1000}
	fmt.Println(account)
}
