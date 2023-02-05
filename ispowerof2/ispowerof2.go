package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	args := os.Args[1]
	var nb int
	nb, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	answer := IsPowerof2(nb)
	fmt.Println(answer)
}

func IsPowerof2(nb int) string {
	i := 2
	if nb == 1 || nb == 2 {
		return "true"
	}
	for {
		if i == nb {
			return "true"
		}
		if i > nb {
			return "false"
		}
		i *= 2
	}
}

// package main
// import (
// 	"os"
// 	"strconv"
// 	"github.com/01-edu/z01"
// )
// func ispowerof2(n int) string {
// 	i := 2
// 	if n == 1 || n == 2 {
// 		return "true"
// 	}
// 	for {
// 		if i == n {
// 			return "true"
// 		}
// 		if i > n {
// 			return "false"
// 		}
// 		i *= 2
// 	}
// }
// func main() {
// 	args := os.Args
// 	var nbr int
// 	if len(args) == 2 {
// 		nbr, _ = strconv.Atoi(args[1])
// 		ans := ispowerof2(nbr)
// 		for _, x := range ans {
// 			z01.PrintRune(x)
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune('\n')
// 	}
// }
