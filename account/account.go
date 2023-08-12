package account

import (
	"errors"
	"fmt"
)

var errCantWithdraw = errors.New("Can't withdraw")

// Account struct
type Account struct {
	owner   string
	balance int
}

// New creates Account
func New(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errCantWithdraw
	}
	a.balance -= amount
	return nil
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account\nhas : ", a.Balance())
}
