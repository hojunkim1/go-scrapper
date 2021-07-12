package main

import (
	"github.com/hojunkim1/learngo/accounts"
)


func main() {
	account := accounts.NewAccount("hojun")
	account.Deposit(10)
}
