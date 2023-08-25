package piscine

func Capitalize(s string) string {
	str := ""
	previousIsAlphanumeric := false

	for index, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if index == 0 || !previousIsAlphanumeric {
				if char >= 'a' && char <= 'z' {
					char -= 32
				}
			} else if char >= 'A' && char <= 'Z' {
				char += 32
			}
			previousIsAlphanumeric = true
		} else {
			previousIsAlphanumeric = false
		}
		str += string(char)
	}

	return str
}
