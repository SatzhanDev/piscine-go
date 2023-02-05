package piscine

func ToLower(s string) string {
	a := []rune(s)
	for i, l := range a {
		if l <= 90 && l >= 65 {
			a[i] = l + 32
		}
	}
	return string(a)
}
