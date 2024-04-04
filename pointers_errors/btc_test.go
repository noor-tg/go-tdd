package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	fmt.Printf("address balance %p\n", &wallet.balance)

	want := 10.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}
