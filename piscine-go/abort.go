package piscine

func Abort(a, b, c, d, e int) int {
	arr := [5]int{a, b, c, d, e}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	medianIndex := 2
	return arr[medianIndex]
}
