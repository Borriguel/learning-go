package pointers

import "fmt"

type Rodocoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Rodocoin
}

func (w *Wallet) Deposit(amount Rodocoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Rodocoin {
	return w.balance
}

func (r Rodocoin) String() string {
	return fmt.Sprintf("%d RDC", r)
}
