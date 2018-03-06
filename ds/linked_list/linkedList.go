package linked_list

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head *Node
}

func InitList() *List {
	return &List{Head: nil}
}

func (list *List) Insert(data interface{}) {
	if list.Head == nil {
		list.Head = &Node{Data: data, Next: nil}
		return
	}

	curr := list.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	// We are at the last element now
	curr.Next = &Node{Data: data, Next: nil}
	return
}

func (list *List) InsertAtHead(data interface{}) {
	if list.Head == nil {
		list.Head = &Node{Data: data, Next: nil}
		return
	}

	temp := &Node{Data: data, Next: list.Head}
	list.Head = temp
}

func (list *List) Reverse() {
	if list.Head == nil {
		fmt.Println("List is empty")
		return
	}

	curr := list.Head
	n := curr.Next

	for n != nil {
		temp := n.Next

		n.Next = curr
		if curr == list.Head {
			curr.Next = nil
		}

		curr = n
		n = temp
	}

	list.Head = curr
}

//func (list *List) ReverseFromIndex()

// Print prints the list to stdout in pretty format
func (list *List) Print() {
	if list.Head == nil {
		return
	}

	curr := list.Head

	for curr.Next != nil {
		fmt.Printf("%v --> ", curr.Data)
		curr = curr.Next
	}

	fmt.Printf("%v\n", curr.Data)
}
