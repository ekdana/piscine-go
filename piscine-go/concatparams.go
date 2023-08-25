package piscine

func ConcatParams(args []string) string {
	var s string = args[0]
	for _, i := range args[1:] {
		s += "\n" + i
	}
	return s
}
