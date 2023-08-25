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

func writeInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	res := ""
	for n > 0 {
		digit := n % 10
		res = string('0'+digit) + res
		n /= 10
	}
	for _, ch := range res {
		z01.PrintRune(ch)
	}
}

func main() {
	points := &point{}

	setPoint(points)
	isX := "x = "
	isY := "y = "
	for _, i := range isX {
		z01.PrintRune(i)
	}

	writeInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	for _, j := range isY {
		z01.PrintRune(j)
	}
	writeInt(points.y)
	z01.PrintRune('\n')
}
