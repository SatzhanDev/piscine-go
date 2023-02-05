package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsPrime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindFactors(nb int) {
	for i := 2; i < nb; i++ {
		if nb%i == 0 && IsPrime(i) == true {
			fmt.Print(i)
			fmt.Print("*")
			FindFactors(nb / i)
			return
		}
	}
	fmt.Println(nb)
}

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		str := os.Args[1]
		nb, err := strconv.Atoi(str)
		if err != nil {
			return
		} else {
			if nb == 0 || nb == 1 {
				return
			}
			if IsPrime(nb) {
				fmt.Println(nb)
				return
			}
			FindFactors((nb))
		}
	}
}
