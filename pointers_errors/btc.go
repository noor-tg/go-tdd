package pointerserrors

import "fmt"

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	fmt.Printf("address balance %p\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	return w.balance
}
