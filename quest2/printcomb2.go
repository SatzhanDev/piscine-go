package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if i < j {
				a1 := i / 10
				b1 := i % 10
				a2 := j / 10
				b2 := j % 10
				z01.PrintRune(rune(a1 + 48))
				z01.PrintRune(rune(b1 + 48))
				z01.PrintRune(' ')
				z01.PrintRune(rune(a2 + 48))
				z01.PrintRune(rune(b2 + 48))
				if i != 98 || j != 99 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
