package piscine

func BasicJoin(elems []string) string {
	str := ""
	for i := range elems {
		str += elems[i]
	}
	return str
}
