package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	} else if n == 0 {
		z01.PrintRune('0')
	}
	slice := []int{}
	for n != 0 {
		slice = append(slice, n%10)
		n = n / 10
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] < 0 {
			slice[i] = -slice[i]
		}
		z01.PrintRune(rune(slice[i] + 48))
	}
}
