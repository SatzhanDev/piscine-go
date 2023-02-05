package main

import (
	"fmt"
	"os"
)

func main() {
	strings := []string{"01", "galaxy", "galaxy 01"}
	args := os.Args[1:]
	alert := "Alert!!!"
	count := 0

	for i := range args {
		for _, v := range strings {
			if args[i] == v {
				count++
			}
		}
	}
	if count >= 1 {
		fmt.Println(alert)
	}
}

// func main() {
// 	word1 := "01"
// 	word2 := "galaxy"
// 	word3 := "galaxy 01"

// 	args := os.Args[1:]

// 	for i := range args {
// 		if args[i] == word1 || args[i] == word2 || args[i] == word3 {
// 			fmt.Println("Alert!!!")
// 			break
// 		}
// 	}
// }
