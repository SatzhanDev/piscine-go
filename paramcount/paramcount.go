package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	} else {
		lenInt := len(os.Args[1:])

		array := make([]rune, lenInt+1)

		for lenInt != 0 {
			array = append(array, rune((lenInt%10)+48))
			lenInt /= 10
		}
		for _, s := range array {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}
