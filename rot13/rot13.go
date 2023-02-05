package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		var answer string
		arg := os.Args[1]
		for i, v := range arg {
			if v >= 97 && v <= 122 || v >= 65 && v <= 90 {
				if v+13 >= 90 && v <= 90 && v >= 65 || v <= 122 && v >= 97 && v+13 >= 122 {
					answer += string(arg[i] - 13)
				} else {
					answer += string(arg[i] + 13)
				}
			} else {
				answer += string(arg[i])
			}
		}
		for _, l := range answer {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}
