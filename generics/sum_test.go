package generics

import "testing"

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
