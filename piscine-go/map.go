package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, arg := range a {
		result[i] = f(arg)
	}
	return result
}
