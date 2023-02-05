package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		// str := os.Args[1]
		// word := ""
		// answer := ""
		// for i := 0; i < len(str); i++ {
		// 	if str[i] != ' ' {
		// 		word += string(str[i])
		// 	} else {
		// 		if word != "" && str[len(str)-1] != ' ' {
		// 			answer = word
		// 			word = ""
		// 		}
		// 	}
		// }
		// if word != "" {
		// 	answer = word
		// 	fmt.Println(answer)
		// } else {
		// 	fmt.Print()
		// }
		str := os.Args[1]
		newStr := ""
		var arrStr []string
		for i, v := range str {
			if v != ' ' && v != rune(str[len(str)-1]) {
				newStr += string(str[i])
			} else if v == ' ' && newStr != "" {
				arrStr = append(arrStr, newStr)
				newStr = ""
			} else if v != ' ' && v == rune(str[len(str)-1]) {
				newStr += string(str[i])
				arrStr = append(arrStr, newStr)
			}
		}
		if len(arrStr) == 0 {
			return
		}
		fmt.Println(arrStr[len(arrStr)-1])
	}
}
