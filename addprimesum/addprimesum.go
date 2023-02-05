package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	} else {
		arg := os.Args[1]

		num, err := strconv.Atoi(arg)
		if err != nil || num <= 0 {
			z01.PrintRune('0')
		}
		sum := 0
		for i := 2; i <= num; i++ {
			if Isprime(i) {
				sum = sum + i
			}
		}
		var revNum []rune
		for sum != 0 {
			revNum = append(revNum, rune(sum%10+48))
			sum = sum / 10
		}
		for i := len(revNum) - 1; i >= 0; i-- {
			z01.PrintRune(revNum[i])
		}
		z01.PrintRune('\n')

	}
}

func Isprime(num int) bool {
	if num < 2 {
		return false
	} else {

		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
}
