package piscine

func StrLen(s string) int {
	a := []rune(s)
	count := 0
	for i := range a {
		i++
		count++
	}
	return count
}
