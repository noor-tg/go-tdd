package generics

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, s string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == s {
			currentBalance -= t.Sum
		}
		if t.To == s {
			currentBalance += t.Sum
		}
		return currentBalance
	}

	return Reduce(transactions, adjustBalance, 0.0)
}
