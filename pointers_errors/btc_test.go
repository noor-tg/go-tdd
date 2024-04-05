package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)

		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw to less than 0 show error", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(45))
		assertBalance(t, wallet, Bitcoin(20))

		assertError(t, err, WithdrawError)
		assertBalance(t, wallet, Bitcoin(20))
	})

}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Should not get Error")
	}
}

func assertError(t testing.TB, got, want error) {
	if got == nil {
		// panic to stop test
		t.Fatal("did not get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
