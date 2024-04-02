package collections

func Sum(input [5]int) int {
	total := 0
	for i := 0; i < 5; i++ {
		total += input[i]
	}
	return total
}
