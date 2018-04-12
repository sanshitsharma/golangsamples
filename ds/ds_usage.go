package main

import (
	"errors"
	"fmt"

	h "github.com/sanshitsharma/golangsamples/ds/heap"
	llist "github.com/sanshitsharma/golangsamples/ds/linked_list"
	"github.com/sanshitsharma/golangsamples/ds/queue"
	"github.com/sanshitsharma/golangsamples/ds/stack"
	"github.com/sanshitsharma/golangsamples/ds/trees"
	"github.com/sanshitsharma/golangsamples/ds/trie"
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
	//stk.Push(2)

	stk.Pop()

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
	size := 3
	q := queue.NewQueue(size)

	nums := []int{10, 20, 30, 40}
	for _, num := range nums {
		if err := q.Enqueue(num); err != nil {
			fmt.Printf("failed to insert: '%v'. err: '%v'\n", num, err)
		}
	}

	q.Print()

	for i := 0; i < size+1; i++ {
		if val, err := q.Dequeue(); err != nil {
			fmt.Printf("failed to dequeue.. err: '%v'\n", err)
		} else {
			fmt.Printf("Dequeued: value = '%v'\n", val)
		}
	}
}

func linkQUsage() {
	q := queue.NewLinkQueue()
	nums := []interface{}{10, 20, 30, 40, "abc"}
	for _, num := range nums {
		q.Enqueue(num)
	}

	q.Print()

	for i := 0; i < 5; i++ {
		if val := q.Dequeue(); val == nil {
			fmt.Printf("failed to dequeue.. queue is empty")
		} else {
			fmt.Printf("Dequeued: value = '%v'\n", val)
		}
	}
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

func trieUsage() {
	t := trie.NewTrie()

	t.Insert(`abcde`)
	t.Insert(`abc`)
	t.Insert(`bdcx`)
	t.Insert(`opqr`)
	t.Insert(`s`)

	str := `s`
	fmt.Printf("isExists('%v') = %v\n", str, t.Search(str))

	str = `sanshit`
	fmt.Printf("isExists('%v') = %v\n", str, t.Search(str))

	str = `abc`
	fmt.Printf("isExists('%v') = %v\n", str, t.Search(str))

	str = `abcd`
	fmt.Printf("isExists('%v') = %v\n", str, t.Search(str))

	str = `abcde`
	fmt.Printf("isExists('%v') = %v\n", str, t.Search(str))
}

func bstUsage() {
	bst := trees.NewBST()

	bst = bst.Insert(10)
	bst = bst.Insert(5)
	bst = bst.Insert(2)
	bst = bst.Insert(7)
	bst = bst.Insert(15)
	bst = bst.Insert(12)
	bst = bst.Insert(17)

	fmt.Println("In order traversal")
	bst.Inorder()
	fmt.Printf("\n\n")

	fmt.Println("Pre order traversal")
	bst.Preorder()
	fmt.Printf("\n\n")

	if node, err := bst.Search(17); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("found", node.Data)
	}

	if node, err := bst.Search(21); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("found", node.Data)
	}

	fmt.Println()
	// Find parent
	values := []int{15, 12, 10, 21}
	for _, trgt := range values {
		if parent, err := bst.Parent(trgt); err != nil {
			fmt.Printf("No parent found for '%v'. error: '%v'\n", trgt, err)
		} else {

			if parent == nil {
				fmt.Printf("'%v' is the root node of tree. No Parent\n", trgt)
			} else {
				fmt.Printf("Parent of '%v' --> '%v'\n", trgt, parent.Data)
			}
		}
	}

	bst, err := bst.Delete(12)
	if err != nil {
		fmt.Println("Failed to delete node")
	}
	fmt.Println("Pre order traversal")
	bst.Preorder()
	fmt.Printf("\n\n")

	fmt.Println("Construct a BST from preorder traversal")
	values = []int{10, 5, 1, 7, 40, 50}
	//values = []int{10, 5, 3, 1}
	//values = []int{10, 40, 50}
	bst2 := trees.ConstructTree(values)

	fmt.Println("Pre order traversal")
	bst2.Preorder()
	fmt.Printf("\n\n")

	fmt.Println("Construct a BST from preorder traversal iteratively")
	values = []int{10, 5, 1, 7, 40, 50}
	//values = []int{10, 5, 3, 1}
	//values = []int{10, 40, 50}
	bst3 := trees.ConstructTreeIter(values)

	fmt.Println("Pre order traversal")
	bst3.Preorder()
	fmt.Printf("\n\n")
}

func main() {
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

	fmt.Println("------------------------ Linked List Usage -------------------------")
	linkedListUsage()

	fmt.Println("\n--------------------------- Heap Usage -----------------------------")
	heapUsage()

	fmt.Println("\n-------------------------- Stack Usage -----------------------------")
	stackUsage()

	// Queue Usage
	fmt.Println("\n------------------------ Array Queue Usage -------------------------")
	qUsage()

	// LinkQueue Usage
	fmt.Println("\n------------------------- List Queue Usage -------------------------")
	linkQUsage()

	// Trie Usage
	fmt.Println("\n---------------------------- TRIE Usage ----------------------------")
	trieUsage()

	// BST Usage
	fmt.Println("\n---------------------------- BST Usage -----------------------------")
	bstUsage()
}
