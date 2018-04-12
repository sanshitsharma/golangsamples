package main

import (
	"errors"
	"fmt"

	"github.com/sanshitsharma/golangsamples/ds/trees"
)

func kthLargestRecurse(node *trees.BSTNode, k, count int) (int, int, error) {
	if node == nil {
		return count, 0, nil
	}

	count, _, _ = kthLargestRecurse(node.Left, k, count)
	count++
	fmt.Println("data:", node.Data, "Curr Count =", count)
	if k == count {
		return count, node.Data, nil
	}
	count, _, _ = kthLargestRecurse(node.Right, k, count)

	return count, 0, errors.New("kth largest does not exist")
}

func kthLargest(node *trees.BSTNode, k int) (int, error) {
	count := 0
	count, elem, err := kthLargestRecurse(node, k, count)

	fmt.Println("Count =", count, "Element =", elem, "Error =", err)
	return elem, err
}

func main() {
	// Create a BST
	bst := trees.NewBST()

	bst = bst.Insert(20)
	bst.Insert(8)
	bst.Insert(22)
	bst.Insert(4)
	bst.Insert(12)
	bst.Insert(10)
	bst.Insert(14)
	bst.Insert(9)

	k := 3
	elem, _ := kthLargest(bst, k)
	fmt.Printf("kth largest element where k = '%v' is '%v'\n", k, elem)
}
