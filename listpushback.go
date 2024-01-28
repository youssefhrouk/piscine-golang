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
	if l.Tail == nil {
		l.Head = s
		l.Tail = s

	} else {
		l.Tail.Next = s
		l.Tail = s
	}
}
