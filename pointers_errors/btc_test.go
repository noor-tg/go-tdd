package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	fmt.Printf("address balance %p\n", &wallet.balance)

	want := Bitcoin(10)
	if got != want {
		t.Errorf("got %d want %s", got, want)
	}

}
