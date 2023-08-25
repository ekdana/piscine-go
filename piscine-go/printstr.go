package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	Byte := []byte(s)

	for _, i := range Byte {
		z01.PrintRune(rune(i))
	}
}
