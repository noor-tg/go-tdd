package generics

import "testing"

func TestReduce(t *testing.T) {
	t.Run("sum two functions by reduce", func(t *testing.T) {
		got := Sum([]int{1, 2, 3})
		want := 6

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Noor",
			To:   "Osama",
			Sum:  100,
		},
		{
			From: "Noor",
			To:   "Ali",
			Sum:  50,
		},
		{
			From: "Noor",
			To:   "Ahmed",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Ali"), 50)
	AssertEqual(t, BalanceFor(transactions, "Ahmed"), 25)
	AssertEqual(t, BalanceFor(transactions, "Osama"), 100)
}
