package main

// import (
// 	"fmt"
// )

func IterativeFactorial(nb int) int {
	// n! = 1 * 2 * ...*n
	result := 1

	if nb == 0 || nb == 1 {
		result = 1
	} else {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	}
	return result
}

// TEST
// func main() {
// 	arg := 3
// 	fmt.Println(IterativeFactorial(arg))
// }
