package piscine

func NRune(s string, n int) rune {
	letter := []rune(s)

	for index, val := range letter {
		if index == n-1 {
			return val
		}
	}
	return 0
}
