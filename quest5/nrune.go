package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	for i, v := range a {
		if i == n-1 {
			return v
		}
	}
	return 0
}
