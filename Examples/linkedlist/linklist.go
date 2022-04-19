package linkedlist

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

// append to the end of the list
func (n *Node) AppendToTail(value int) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = NewNode(value)
}

// print the list
func (n *Node) Print() {
	for n != nil {
		print(n.Value, " ")
		n = n.Next
	}
	println()
}

// create a linked list
func CreateLinkedList(values []int) *Node {
	head := NewNode(values[0])
	for i := 1; i < len(values); i++ {
		head.AppendToTail(values[i])
	}
	return head
}

type LinkedList struct {
	Head *Node
}

// append to the end of the list
func (l *LinkedList) AppendToTail(value int) {
	if l.Head == nil {
		l.Head = NewNode(value)
		return
	}
	n := l.Head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = NewNode(value)
}


// delete the given node
func (l *LinkedList) DeleteNode(n *Node) {
	if l.Head == n {
		l.Head = l.Head.Next
		return
	}
	prev := l.Head
	for prev.Next != nil {
		if prev.Next == n {
			prev.Next = prev.Next.Next
			return
		}
		prev = prev.Next
	}
}

// print the list
func (l *LinkedList) Print() {
	n := l.Head
	for n != nil {
		print(n.Value, " ")
		n = n.Next
	}
	println()
}
