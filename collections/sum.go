package collections

func Sum(input []int) int {
	total := 0
	for _, number := range input {
		total += number
	}
	return total
}
