package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, strings := range a {
		toRune(strings)
	}
}

func toRune(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
