package piscine

func BasicAtoi2(s string) int {
	num := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			num = num*10 + digit
		} else {
			return 0
		}
	}
	return num
}
