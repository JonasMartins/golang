package pointers

import (
    "fmt"
    "errors"
)

type Bitcoin int
var ErrInsufficientFunds = errors.New("cannot withdraw, insifficient funds")

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return ErrInsifficientFounds
    }
    w.balance -= amount
    return nil
}


