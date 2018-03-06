package stack

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func newNode(val interface{}) *Node {
	return &Node{value: val, prev: nil, next: nil}
}