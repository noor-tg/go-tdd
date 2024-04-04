package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		if got == nil {
			t.Fatal("did not get an error but wanted one")
		}
		if got.Error() != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(15))

	})

	t.Run("withdraw to less than 0 show error", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(45))

		assertBalance(t, wallet, Bitcoin(20))

		assertError(t, err, WithdrawErrorMessage)
		assertBalance(t, wallet, Bitcoin(20))
	})

}
