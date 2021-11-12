package main

import (
	"fmt"

	"github.com/kojh0111/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jh")
	fmt.Println(account)
}
