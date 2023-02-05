package piscine

func IsUpper(s string) bool {
	a := []rune(s)
	for _, v := range a {
		if v <= 'A' || v >= 'Z' {
			return false
		}
	}
	return true
}
