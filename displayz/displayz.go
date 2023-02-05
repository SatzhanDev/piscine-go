package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 2 {
		z01.PrintRune('z')
		z01.PrintRune('\n')
	} else {
		for i, v := range os.Args[1] {
			if v == 'z' || v == 'Z' {
				z01.PrintRune(rune(os.Args[1][i]))
			} else {
				z01.PrintRune('z')
				break
			}
		}
		z01.PrintRune('\n')
	}
}
