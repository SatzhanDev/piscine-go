package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		arg := []rune(os.Args[1])
		for i, v := range arg {
			if v <= 'z' && v >= 'a' {
				arg[i] = v - 32
			} else if v <= 'Z' && v >= 'A' {
				arg[i] = v + 32
			}
		}
		fmt.Println(string(arg))
	}
}
