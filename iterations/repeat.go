package iterations

func Repeat(a string) string {
	var str string
	for i := 0; i < 5; i++ {
		str = str + a
	}
	return str
}
