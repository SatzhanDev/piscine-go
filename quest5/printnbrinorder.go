package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	array := []int{}

	for n != 0 {
		ostatok := n % 10
		array = append(array, ostatok)
		n = n / 10
	}
	len := 0
	for range array {
		len++
	}
	for k := 1; k < len; k++ {
		for i := 1; i < len; i++ {
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
			}
		}
	}
	for _, nb := range array {
		z01.PrintRune(rune(nb + 48))
	}
}
