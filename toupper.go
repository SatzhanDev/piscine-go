package piscine

func ToUpper(s string) string {
	a := []rune(s)
	for i, l := range a {
		if l <= 122 && l >= 97 {
			a[i] = l - 32
		}
	}
	return string(a)
}
