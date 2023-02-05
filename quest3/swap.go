package piscine

func Swap(a *int, b *int) {
	var p int
	var l int
	p = *a
	l = *b
	*a = l
	*b = p
}
