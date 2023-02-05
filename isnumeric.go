package piscine

func IsNumeric(s string) bool {
	a := []rune(s)
	for _, v := range a {
		if v > 57 || v < 48 {
			return false
		}
	}
	return true
}
