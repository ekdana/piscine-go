package piscine

type NodeLI struct {
	Data interface{}
	Next *NodeLI
}

func ListAt(l *NodeLI, pos int) *NodeLI {
	index := 0
	count := 0
	head := l

	for head != nil {
		index++
		head = head.Next
	}
	if pos <= index {
		for l != nil {
			if count == pos {
				return l
			}
			count++
			l = l.Next
		}
	}
	return nil
}
