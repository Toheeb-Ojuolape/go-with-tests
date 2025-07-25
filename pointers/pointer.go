package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type String interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance > amount {
		w.balance -= amount
		return nil
	}
	return ErrInsufficientFunds
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
