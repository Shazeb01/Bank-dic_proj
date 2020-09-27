package main

import (
	"fmt"

	"github.com/Shazeb01/Bank-dic_proj/pkg/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(100)
	fmt.Println(account)

}
