package generics

// use two types of generics to convert some collection to single value
func Reduce[A, B any](arr []A, acculmator func(old B, New A) B, init B) B {
	var total = init

	for _, el := range arr {
		total = acculmator(total, el)
	}

	return total
}

func Sum(input []int) int {
	add := func(a, b int) int { return a + b }

	return Reduce(input, add, 0)
}
