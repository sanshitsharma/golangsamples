package main

import (
	"github.com/sanshitsharma/golangsamples/ds/linked_list"
	"github.com/sanshitsharma/golangsamples/ds/trees"
)

func bstFromSortedListRecurse(lNode **linked_list.Node, n int) *trees.BSTNode {
	if n <= 0 {
		return nil
	}

	// Recurse left until you can
	left := bstFromSortedListRecurse(lNode, n/2)

	// Allocate memory for root and attach left sub-tree
	data, _ := (*lNode).Data.(int)
	root := &trees.BSTNode{Data: data, Left: nil, Right: nil}
	root.Left = left

	// Advance the pointer because not root has been used
	*lNode = (*lNode).Next

	// Now recurse right by halving the right. The right half will be equal to
	// n - n/2 - 1. The n/2 left nodes and 1 root node
	root.Right = bstFromSortedListRecurse(lNode, (n - n/2 - 1))

	return root
}

func bstFromSortedList(lst *linked_list.List) *trees.BSTNode {
	n := lst.Count()
	return bstFromSortedListRecurse(&lst.Head, n)
}

func main() {
	// Create a linked list
	lst := linked_list.InitList()
	lst.Insert(1)
	lst.Insert(2)
	lst.Insert(3)
	lst.Insert(4)
	lst.Insert(5)
	lst.Insert(6)
	lst.Insert(7)

	//lst.Print()
	bst := bstFromSortedList(lst)
	bst.Preorder()
}
