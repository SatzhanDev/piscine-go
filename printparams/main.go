package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, words := range arg[1:] {
		for _, letter := range words {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
