package piscine

type Nodell struct {
	Data interface{}
	Next *NodeL
}

type Listll struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *Listll, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}

	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}
