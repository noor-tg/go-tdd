package generics

// use two types of generics to convert some collection to single value
func Reduce[A, B any](arr []A, acculmator func(old B, New A) B, init B) B {
	var total = init

	for _, el := range arr {
		total = acculmator(total, el)
	}

	return total
}

func Find[A any](items []A, predate func(p A) bool) (A, bool) {
	var def A
	for _, n := range items {
		if predate(n) {
			return n, true
		}
	}
	return def, false
}
