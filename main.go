package main

import (
	"fmt"

	"github.com/Shazeb01/Bank-dic_proj/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())

}
