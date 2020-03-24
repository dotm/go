package main

import (
	"errors"
	"fmt"
)

func main() {}

//ErrInsufficientFunds is used when user try withdrawing Bitcoin with insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Bitcoin represent the number of bitcoin in a wallet
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet represents a Bitcoin wallet
type Wallet struct {
	balance Bitcoin
}

//Deposit adds the amount of money in a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance returns the total amount of money in a wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Withdraw substracts the amount of money in a wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
