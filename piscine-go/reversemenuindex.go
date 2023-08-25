package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)

	reversedMenu := make([]string, length)

	for i := 0; i < length; i++ {
		reversedMenu[length-1-i] = menu[i]
	}

	return reversedMenu
}
