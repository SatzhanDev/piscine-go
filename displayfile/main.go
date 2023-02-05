package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	length := len(args)
	if length == 0 {
		fmt.Println("File name missing")
		return
	} else if length > 1 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := args[0]

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println()
		return
	}
	fmt.Print(string(content))
}
