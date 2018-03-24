package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	top           *Node
	middle        *Node
	midUpdCounter int
	elemCounter   int
}

func NewStack() *Stack {
	return &Stack{top: nil, middle: nil, midUpdCounter: 0}
}

func (stack *Stack) Push(value interface{}) {
	if stack.top == nil {
		stack.top = newNode(value)
		stack.middle = stack.top
		stack.elemCounter += 1
		return
	}

	//fmt.Println("Pushing.. ", value, " curr mid:", stack.middle.value, " mid counter:", stack.midUpdCounter)
	// Create a new node
	node := newNode(value)

	// next and prev pointer assignments
	stack.top.next = node
	node.prev = stack.top

	// Now set the top to new node
	stack.top = node
	stack.midUpdCounter += 1
	stack.elemCounter += 1

	if stack.midUpdCounter == 2 {
		//fmt.Println("modifying mid after push.. curr mid:", stack.middle.value)
		// Modify mid
		stack.middle = stack.middle.next
		stack.midUpdCounter = 0
	}
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.top == nil {
		return nil, errors.New("empty stack")
	}

	//fmt.Println("Popping.. curr mid:", stack.middle.value, " mid counter:", stack.midUpdCounter)
	val := stack.top.value
	temp := stack.top.prev
	if temp != nil {
		temp.next = nil
	}
	stack.top.prev = nil
	stack.top = temp

	stack.elemCounter += 1
	stack.midUpdCounter -= 1

	if stack.midUpdCounter == -2 {
		// Modify mid
		stack.middle = stack.middle.prev
		stack.midUpdCounter = 0
	}

	return val, nil
}

// Peek returns the top element of the stack without removing
// it from the stack
func (stack *Stack) Peek() (interface{}, error) {
	if stack.top == nil {
		return nil, errors.New("empty stack")
	}

	return stack.top.value, nil
}

// IsEmpty checks if the stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}

func (stack *Stack) FindMiddle() (interface{}, error) {
	if stack.top == nil {
		return nil, errors.New(`empty stack`)
	}

	return stack.middle.value, nil
}

/*
func (stack *Stack) DeleteMiddle() error {
	if stack.top == nil {
		return errors.New(`empty stack`)
	}

	if stack.elemCounter%2 == 0 {
		// point stack.middle to next
		stack.middle = stack.middle.next
		stack.middle.prev.prev.next = stack.middle
		stack.middle.prev = stack.middle.prev.prev
	} else {
		stack.middle = stack.middle.prev
		stack.middle.next.next.prev =
	}
}
*/

func (stack *Stack) Print() {
	curr := stack.top
	for curr != nil {
		fmt.Printf("%v ", curr.value)
		curr = curr.prev
	}
	fmt.Println()
}
