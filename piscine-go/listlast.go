package piscine

type NodeLL struct {
	Data interface{}
	Next *NodeL
}

type ListLL struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	i := l.Head

	for i.Next != nil {
		i = i.Next
	}
	return i.Data
}
