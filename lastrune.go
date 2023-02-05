package piscine

func LastRune(s string) rune {
	a := []rune(s)
	return a[len(s)-1]
}
