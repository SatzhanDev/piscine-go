package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	strX := "x = "
	strY := "y = "

	arrayX := []rune{}
	arrayY := []rune{}

	valueX := points.x
	valueY := points.y

	for valueX != 0 {
		arrayX = append(arrayX, rune(valueX%10))
		valueX /= 10
	}

	for _, v := range strX {
		z01.PrintRune(v)
	}

	for i := len(arrayX) - 1; i >= 0; i-- {
		z01.PrintRune(arrayX[i] + 48)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')

	for valueY != 0 {
		arrayY = append(arrayY, rune(valueY%10))
		valueY /= 10
	}

	for _, v := range strY {
		z01.PrintRune(v)
	}

	for i := len(arrayY) - 1; i >= 0; i-- {
		z01.PrintRune(arrayY[i] + 48)
	}
	z01.PrintRune('\n')
}
