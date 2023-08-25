package piscine

func StrLen(s string) int {
	count := 0
	for i := range []rune(s) {
		count = i + 1
	}

	return count
}
