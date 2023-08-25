package main

import "github.com/01-edu/z01"

func printLine(ch1, ch2, ch3 rune, width int) {
	z01.PrintRune(ch1)

	for i := 0; i < width-2; i++ {
		z01.PrintRune(ch2)
	}

	if width > 1 {
		z01.PrintRune(ch3)
	}

	z01.PrintRune('\n')
}

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	printLine('o', '-', 'o', x)

	for i := 0; i < y-2; i++ {
		printLine('|', ' ', '|', x)
	}

	if y > 1 {
		printLine('o', '-', 'o', x)
	}
}

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	printLine('/', '*', '\\', x)

	for i := 0; i < y-2; i++ {
		printLine('*', ' ', '*', x)
	}

	if y > 1 {
		printLine('\\', '*', '/', x)
	}
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	printLine('A', 'B', 'A', x)

	for i := 0; i < y-2; i++ {
		printLine('B', ' ', 'B', x)
	}

	if y > 1 {
		printLine('C', 'B', 'C', x)
	}
}

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	printLine('A', 'B', 'C', x)

	for i := 0; i < y-2; i++ {
		printLine('B', ' ', 'B', x)
	}

	if y > 1 {
		printLine('A', 'B', 'C', x)
	}
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	printLine('A', 'B', 'C', x)

	for i := 0; i < y-2; i++ {
		printLine('B', ' ', 'B', x)
	}

	if y > 1 {
		printLine('C', 'B', 'A', x)
	}
}

func main() {
	// QuadA(5, 3)
	// QuadA(5, 1)
	// QuadA(1, 1)
	// QuadA(1, 5)

	// QuadC(5, 3)
	// QuadC(5, 1)
	// QuadC(1, 1)
	// QuadC(1, 5)

	// QuadD(5, 3)
	// QuadD(5, 1)
	// QuadD(1, 1)
	// QuadD(1, 5)

	QuadE(5, 3)
	QuadE(5, 1)
	QuadE(1, 1)
	QuadE(1, 5)
}
