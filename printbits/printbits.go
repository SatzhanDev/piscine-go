package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		number, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("00000000")
		} else {
			var array []int
			for i := 0; i < 8; i++ {
				array = append(array, number%2)
				number = number / 2
			}
			for i := 7; i >= 0; i-- {
				z01.PrintRune(rune(array[i] + 48))
			}
		}
	}
}
