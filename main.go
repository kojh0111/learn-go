package main

import (
	"fmt"

	"github.com/kojh0111/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jh")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
