package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= '9'; i++ {
		for j := 0; j <= '9'; j++ {
			w := j + 1
			for k := i; k <= '9'; k++ {
				for x := w; x <= '9'; x++ {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))

					z01.PrintRune(' ')
					z01.PrintRune(rune(k))
					z01.PrintRune(rune(x))
					if !(i == '9' && j == '8' && k == '9' && x == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				w = '0'
			}
		}
		z01.PrintRune('\n')
	}
}
