package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		word := os.Args[1]

		pigWord := ""
		pigLastlet := ""

		if IsVowel(string(word[0])) {
			pigWord += string(word)
			pigWord += "ay"
		} else {
			for i := 0; i < len(word); i++ {
				if IsVowel(string(word[i])) == false {
					pigLastlet += string(word[i])
				} else {
					pigWord = string(word[i:]) + pigLastlet + "ay"
					break
				}
			}
		}
		if pigWord != "" {
			fmt.Println(pigWord)
		} else {
			fmt.Println("No vowels")
		}
	}
}

func IsVowel(a string) bool {
	arrVowel := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for i := range arrVowel {
		if a == arrVowel[i] {
			return true
		}
	}
	return false
}
