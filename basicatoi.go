package piscine

func BasicAtoi(s string) int {
	a := []byte(s)
	result := 0
	pow := 1
	for i := len(a) - 1; i >= 0; i-- {
		result += ((int(a[i]) - 48) * pow)
		pow *= 10
	}
	return result
}
