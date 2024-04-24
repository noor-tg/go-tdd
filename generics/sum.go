package generics

func Reduce[A any](arr []A, acculmator func(old A, New A) A, init A) A {
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
