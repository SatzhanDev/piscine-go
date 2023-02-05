package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	for i, a := range arg {
		if i > 1 {
			z01.PrintRune(a)
		}
	}
	z01.PrintRune('\n')
}
