package utils

func Repeat(c string, t int) string {
	var str string

	for range t {
		str += c
	}

	return str
}
