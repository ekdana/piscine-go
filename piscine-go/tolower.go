package piscine

func ToLower(s string) string {
	str := ""
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			i = i + 32
		}
		str += string(rune(i))
	}
	return str
}
