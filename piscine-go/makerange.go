package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	var s []int = make([]int, max-min)
	for i, j := min, 0; i < max; i++ {
		s[j] = i
		j++
	}

	return s
}
