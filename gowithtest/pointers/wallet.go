package pointers

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Bitcoint int

func (b Bitcoint) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoint
}

func (w *Wallet) Deposit(amount Bitcoint) {
	w.balance = amount
}

func (w *Wallet) Balance() Bitcoint {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoint) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
