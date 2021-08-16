package main

import (
	"fmt"
	accounts "github/whguddnr9/learngo/acoounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
