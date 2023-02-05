package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		arg := os.Args[1]
		newArg := ""
		var arrOfstr []string

		for i, v := range arg {
			if v != ' ' && v != '\t' {
				newArg += string(arg[i])
			} else if (v == ' ' || v == '\t') && newArg != "" {
				arrOfstr = append(arrOfstr, newArg)
				newArg = ""
			}
		}
		if newArg != "" {
			arrOfstr = append(arrOfstr, newArg)
		}

		for j, v := range arrOfstr {
			if j != len(arrOfstr)-1 {
				for i := range v {
					z01.PrintRune(rune(v[i]))
				}
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			} else {
				for i := range v {
					z01.PrintRune(rune(v[i]))
				}
				z01.PrintRune('\n')
			}
		}
	}
}
