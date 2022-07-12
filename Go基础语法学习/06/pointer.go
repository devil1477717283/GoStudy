package _6

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Stringer interface {
	String() string
	//String
}

type Wallet struct {
	balance Bitcoin //Go 中，如果一个符号（例如变量、类型、函数等）是以小写符号开头，那么它在 定义它的包之外 就是私有的。
}

func (w *Wallet) Deposit(v Bitcoin) {
	w.balance += v
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(v Bitcoin) error {
	if w.balance-v < 0 {
		return InsufficientFundsError
	}
	w.balance -= v
	return nil
}
func Hello() string {
	return "Hello,World"
}
