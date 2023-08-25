package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, ch := range s {
		digit := int(ch - '0')
		num = num*10 + digit
	}
	return num
}
