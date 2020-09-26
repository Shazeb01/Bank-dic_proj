package accounts

import "errors"

//Account struct
type Account struct {
	owner   string
	balance int
}

var errNomoney = errors.New("Can't withdraw")

//NewAccount Test
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit amount
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance amount
func (a Account) Balance() int {
	return a.balance
}

//Withdraw amount
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNomoney
	}
	a.balance -= amount
	return nil
}
