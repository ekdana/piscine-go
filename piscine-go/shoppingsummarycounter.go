package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counter := make(map[string]int)
	items := splitString(str, " ")

	for _, item := range items {
		counter[item]++
	}

	return counter
}

func splitString(s string, sep string) []string {
	result := make([]string, 0)
	start := 0

	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == sep[0] {
			result = append(result, s[start:i])
			start = i + 1
		}
	}

	return result
}
