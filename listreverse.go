package piscine

func ListReverse(l *List) {
	c := l.Head
	p := l.Head
	p = nil
	for c != nil {
		next := c.Next
		c.Next = p
		p = c
		c = next
	}
	temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
}
