package piscine

func LastRune(s string) rune {
	lastLetter := []rune(s)
	length := 0

	for i := range s {
		length = i + 1
	}
	return lastLetter[length-1]
}
