package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	s := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = s
		l.Tail = s
	}
	l.Tail.Next = s
	l.Tail = s
}
