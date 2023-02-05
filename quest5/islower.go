package piscine

func IsLower(s string) bool {
	a := []rune(s)
	for _, v := range a {
		if v < 97 || v > 122 {
			return false
		}
	}
	return true
}
