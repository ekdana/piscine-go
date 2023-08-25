package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	bubbleSort(args)
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func bubbleSort(arr []string) {
	ln := len(arr)
	for i := 0; i < ln-1; i++ {
		for j := 0; j < ln-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
