package piscine

func AlphaCount(s string) int {
	length := 0

	for _, i := range s {
		if (i <= 90 && i >= 65) || (i <= 122 && i >= 97) {
			length = length + 1
		}
	}
	return length
}
