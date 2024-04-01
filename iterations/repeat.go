package iterations

func Repeat(a string) string {
	var str string
	for range 5 {
		str += a
	}
	return str
}
