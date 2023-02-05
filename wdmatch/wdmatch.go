package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	} else {
		var answer string
		// count := 0
		arg1 := os.Args[1]
		arg2 := os.Args[2]

		// for i := 0; i < len(arg1); i++ {
		// 	for j := count; j < len(arg2); j++ {
		// 		if arg1[i] == arg2[j] {
		// 			answer += string(arg1[i])
		// 			break
		// 		}
		// 		count++
		// 	}
		// }
		for i := range arg2 {
			for j := range arg1 {
				if arg1[j] == arg2[i] {
					answer += string(arg1[j])
					break
				}
			}
		}
		if arg1 == answer {
			for _, v := range answer {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		}
	}
}
