package generics

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, s string) float64 {
	balance := 0.0

	for _, t := range transactions {
		if t.From == s {
			balance -= t.Sum
		}
		if t.To == s {
			balance += t.Sum
		}
	}

	return balance
}
