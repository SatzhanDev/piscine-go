package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println()
		return
	}
	arg := os.Args[1]
	word := ""
	var array []string
	for _, v := range arg {
		if v != ' ' {
			word += string(v)
		} else {
			array = append(array, word)
			word = ""
		}
	}

	if word != "" {
		array = append(array, word)
	}

	for i := len(array) - 1; i >= 0; i-- {
		fmt.Print(array[i])
		if array[i] != array[0] {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
