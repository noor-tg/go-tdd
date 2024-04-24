package generics

func Sum(input []int) int {
	add := func(a, b int) int { return a + b }

	return Reduce(input, add, 0)
}
