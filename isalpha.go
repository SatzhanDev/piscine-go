package piscine

func IsAlpha(s string) bool {
	a := []rune(s)
	for _, v := range a {
		if v >= 0 && v < 48 || v > 57 && v < 65 || v > 90 && v < 97 || v > 122 && v <= 127 {
			return false
		}
	}
	return true
}
