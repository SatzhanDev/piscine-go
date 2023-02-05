package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	} else {
		nb, err := strconv.Atoi(arg[0])
		if err != nil {
			return
		} else {
			if nb < 0 {
				return
			}
			for i := 1; i < 10; i++ {
				mult(i)
				z01.PrintRune(' ')
				z01.PrintRune('x')
				z01.PrintRune(' ')
				mult(nb)
				z01.PrintRune(' ')
				z01.PrintRune('=')
				z01.PrintRune(' ')
				mult(i * nb)
				z01.PrintRune(10)
			}
		}
	}
}

func mult(nb int) {
	if nb >= 10 {
		mult(nb / 10)
	}
	z01.PrintRune(rune(nb%10 + 48))
}
