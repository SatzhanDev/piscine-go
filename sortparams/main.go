package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i, word1 := range arg {
		for j, word2 := range arg {
			if word1 < word2 {
				arg[j], arg[i] = arg[i], arg[j]
			}
		}
	}
	for _, words := range arg {
		for _, letters := range words {
			z01.PrintRune(letters)
		}
		z01.PrintRune('\n')

	}
}
