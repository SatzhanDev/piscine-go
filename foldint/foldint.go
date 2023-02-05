package main

import "fmt"

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, v := range a {
		n = f(n, v)
	}
	fmt.Println(n)
}

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

func Sub(a int, b int) int {
	return a - b
}
