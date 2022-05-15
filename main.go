package main

import (
	"fmt"

	"github.com/dddolmaeng/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("dddolmaeng")
	account.Deposit(10)
	fmt.Println(account)
}