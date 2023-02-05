package piscine

func IsPrintable(s string) bool {
	a := []rune(s)
	for _, v := range a {
		if v >= 0 && v <= 31 {
			return false
		}
	}
	return true
}
