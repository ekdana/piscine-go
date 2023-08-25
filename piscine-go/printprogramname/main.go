package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	lastSlashIndex := -1
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			lastSlashIndex = i
			break
		}
	}
	baseName := name[lastSlashIndex+1:]

	for _, char := range baseName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
