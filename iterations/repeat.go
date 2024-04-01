package iterations

const repeatCount = 5

func Repeat(a string) string {
	var str string
	for range repeatCount {
		str += a
	}
	return str
}
