package main

import "fmt"

import "errors"

type Bitcoin int

var InsufficientFundsError = errors.New("can not withdraw, insufficient founds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	ballence Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.ballence += amount
}

func (w Wallet) Ballence() Bitcoin {
	return w.ballence
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.ballence < amount {
		return InsufficientFundsError
	}
	w.ballence -= amount
	return nil
}
