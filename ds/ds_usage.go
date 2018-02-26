package main

import (
	"fmt"

	h "github.com/sanshitsharma/golangsamples/ds/heap"
	llist "github.com/sanshitsharma/golangsamples/ds/linked_list"
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

func main() {
	//linkedListUsage()
	heapUsage()
}
