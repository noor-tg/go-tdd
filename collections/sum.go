package collections

func Sum(input []int) int {
	total := 0
	for _, number := range input {
		total += number
	}
	return total
}

// accept dynamic input or int slices
// return total of each numbers collection as slice
func SumAll(input ...[]int) []int {
	// init empty slice
	var sums []int
	for _, collection := range input {
		// make new slices from old one with item
		sums = append(sums, Sum(collection))
	}

	return sums
}

func SumAllTails(input ...[]int) []int {
	// init empty slice
	var sums []int
	for _, collection := range input {
		tail := collection[1:]
		// make new slices from old one with item
		sums = append(sums, Sum(tail))
	}

	return sums
}

func SumAllHeadNick(input ...[]int) []int {
	// init empty slice
	var sums []int
	for _, collection := range input {
		head := collection[0:2]
		// make new slices from old one with item
		sums = append(sums, Sum(head))
	}

	return sums
}
