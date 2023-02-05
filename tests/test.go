package main

import (
	"fmt"
)

func main() {
	var x int = 5
	var y *int = &x
	*y = 10
	fmt.Println(x, y)
	print(x)
	print("\n")
}
