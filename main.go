package main

import (
	"fmt"

	"github.com/kojh0111/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jh")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
