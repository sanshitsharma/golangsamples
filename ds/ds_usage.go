package main

import (
	"errors"
	"fmt"

	h "github.com/sanshitsharma/golangsamples/ds/heap"
	llist "github.com/sanshitsharma/golangsamples/ds/linked_list"
	"github.com/sanshitsharma/golangsamples/ds/queue"
	"github.com/sanshitsharma/golangsamples/ds/stack"
)

func linkedListUsage() {
	fmt.Println("Testing linkedList")
	list := llist.InitList()

	list.Insert(5)
	list.Insert(4)
	list.Insert(3)
	list.Insert(2)
	list.Insert(1)
	list.Print()

	list.Reverse()
	list.Print()
}

func heapUsage() {
	fmt.Println("Testing heaps")
	heap, err := h.NewHeap(h.Max)
	if err != nil {
		fmt.Println("failed to create heap")
	}

	elems := []int{3, 1, 2, 8}
	for _, elem := range elems {
		heap.Insert(elem)
	}
	heap.Print()

	//heap.Pop()
	heap.Delete(8)
	heap.Print()

	// Deletion of heap
	/*
		if err := heap.Delete(6); err != nil {
			fmt.Println(err)
		} else {
			heap.Print()
		}
	*/
}

func stackUsage() {
	fmt.Println("Testing stacks")
	stk := stack.NewStack()

	stk.Push(1)
	stk.Push(2)

	fmt.Println(stk.FindMiddle())

	/*
		stk.Push(3)
		stk.Push(4)
		stk.Push(5)

		stk.Print()

		fmt.Println(stk.FindMiddle())

		stk.Pop()
		fmt.Println(stk.FindMiddle())

		stk.Push(6)
		fmt.Println(stk.FindMiddle())

		stk.Pop()
		fmt.Println(stk.FindMiddle())
		stk.Print()

		stk.Push(7)
		stk.Push(8)
		stk.Print()
		fmt.Println(stk.FindMiddle())
	*/
}

func qUsage() {
	q := queue.NewQueue(3)

	if err := q.Enqueue(10); err != nil {
		fmt.Println(err)
	}
	if err := q.Enqueue(20); err != nil {
		fmt.Println(err)
	}
	if err := q.Enqueue(30); err != nil {
		fmt.Println(err)
	}
	if err := q.Enqueue(40); err != nil {
		fmt.Println(err)
	}

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

func swapInPairs(list *llist.List) error {
	if list == nil {
		return errors.New("invalid list")
	}

	if list.Head == nil {
		fmt.Println("empty list. nothing to swap")
		return nil
	}

	// No the magic
	prev := list.Head
	left := list.Head
	curr := left.Next

	for curr != nil {
		temp := curr.Next
		curr.Next = left
		left.Next = temp
		if prev == list.Head {
			list.Head = curr
		} else {
			prev.Next = curr
		}

		if temp == nil {
			break
		}

		prev = left
		left = temp
		curr = temp.Next
	}

	return nil
}
func main() {
	//linkedListUsage()
	//heapUsage()
	//stackUsage()

	/*
		// Given a linkedlist, swap the nodes in pairs
		list := llist.InitList()
		list.Insert(1)
		list.Insert(2)
		list.Insert(3)
		list.Insert(4)
		list.Insert(5)

		list.Print()

		if err := swapInPairs(list); err != nil {
			fmt.Println("failed to swap nodes in pairs:", err)
		}

		list.Print()
	*/

	// Queue Usage
	qUsage()
}
