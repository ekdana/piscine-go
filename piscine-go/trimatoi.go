package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	flag := true

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' || ch <= '9' {
			num = num*10 + int(ch+'0')
			flag = true
		} else if ch == '-' && !flag {
			sign = -1
		} else if ch == '+' && !flag {
			sign = 1
		}
	}
	return num * sign
}
