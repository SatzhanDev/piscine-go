package main

import "fmt"

func main() {
	str := "HelloHALhowHALareHALyou?"
	fmt.Println(Split(str, "HAL"))
}

func Split(str, charset string) []string {
	sublength := len(charset)
	for i := 0; i < sublength; i++ {
		str += " "
	}
	word := ""
	arr := []string{}
	for i := 0; i < len(str); i++ {
		if str[i:i+sublength] != charset && str[i] != str[len(str)-sublength] {
			word += string(str[i])
		} else {
			if word != "" {
				arr = append(arr, word)
				word = ""
				if str[i] == ' ' {
					break
				}
			}
		}
		if str[i:i+sublength] == charset {
			i = i + sublength - 1
		}
	}
	if word != "" {
		arr = append(arr, word)
	}
	return arr
}
