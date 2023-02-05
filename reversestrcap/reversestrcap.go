package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Capitalize(s string) string {
	a := []rune(s)
	for index, letter := range a {
		if check(letter) {
			if index == 0 || check(a[index-1]) == false {
				if letter >= 'a' && letter <= 'z' {
					a[index] = letter - 32
				}
			} else {
				if letter >= 'A' && letter <= 'Z' {
					a[index] = letter + 32
				}
			}
		}
	}
	return string(a)
}

func check(r rune) bool {
	if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}

func main() {
	if len(os.Args) <= 1 {
		return
	}
	args := os.Args[1:]
	for _, str := range args {
		var revresedString []byte
		for i := range str {
			revresedString = append(revresedString, str[len(str)-1-i])
		}
		capString := Capitalize(string(revresedString))
		for i := len(capString) - 1; i >= 0; i-- {
			z01.PrintRune(rune(capString[i]))
		}
		z01.PrintRune('\n')
	}
}

// func Capitalize(s string) string {
// 	a := 0
// 	array := []rune(s)
// 	for range s {
// 		a++
// 	}
// 	for i := 0; i < a; i++ {
// 		if array[i] >= 'A' && array[i] <= 'Z' {
// 			array[i] = array[i] + 32
// 		}
// 		if array[0] >= 'a' && array[0] <= 'z' {
// 			array[0] = array[0] - 32
// 		}
// 		if i > 0 {
// 			if array[i-1] > 'Z' && array[i-1] < 'a' || array[i-1] < 'A' && array[i-1] > '9' || array[i-1] < '0' && array[i-1] != '\'' || array[i-1] > 'z' {
// 				if array[i] >= 'a' && array[i] <= 'z' {
// 					array[i] = array[i] - 32
// 				}
// 			}
// 		}
// 	}
// 	return string(array)
// }

// package main

// import "os"

// func main() {
// 	if len(os.Args) <= 1 {
// 		return
// 	} else {
// 		args := os.Args[1:]
// 		lengthArr := len(args)
// 		str1 := ""
// 		str2 := ""
// 		str3 := ""
// 		for i := 0; i <= lengthArr - 1; i++{
// 				str1 = arg[i]
// 			}
// 		}
