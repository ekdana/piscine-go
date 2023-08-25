package main

import (
	"fmt"
	"os"
)

func checkAlert(args []string) bool {
	keywords := []string{"01", "galaxy", "galaxy 01"}

	for _, arg := range args {
		for _, keyword := range keywords {
			if arg == keyword {
				return true
			}
		}
	}

	return false
}

func main() {
	args := os.Args[1:]
	if checkAlert(args) {
		fmt.Println("Alert!!!")
	} else {
		fmt.Print()
	}
}
