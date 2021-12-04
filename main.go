package main

import (
	"fmt"

	"github.com/Tozinoo/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("hyunseok")
	account.Deposit(10)
	fmt.Println(account)

}
