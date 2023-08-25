package piscine

type NodeLi struct {
	Data interface{}
	Next *NodeL
}

type Listi struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *Listi, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	} else {
		n.Next = l.Head
		l.Head = n
	}
}
