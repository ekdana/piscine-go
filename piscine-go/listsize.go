package piscine

type Nodes struct {
	Data interface{}
	Next *NodeL
}

type Lists struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *Lists) int {
	size := 0
	for l.Head != nil {
		l.Head = l.Head.Next
		size++
	}
	return size
}
