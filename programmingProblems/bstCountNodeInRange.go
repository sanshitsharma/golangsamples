package main

import (
	"fmt"

	"github.com/sanshitsharma/golangsamples/ds/trees"
)

/*
Given a Binary Search Tree (BST) and a range, count the number of nodes in the BST that lie in the given range.
Example:
Input
10 5 50 1 40 100
5 45

Output
3
*/

func getCountOfNode(root *trees.BSTNode, low, high int) int {
	if root == nil {
		return 0
	}

	if root.Data >= low && root.Data <= high {
		return 1 + getCountOfNode(root.Left, low, high) + getCountOfNode(root.Right, low, high)
	} else if root.Data > high {
		return getCountOfNode(root.Left, low, high)
	} else {
		return getCountOfNode(root.Right, low, high)
	}
}

func main() {
	bst := trees.NewBST()
	bst = bst.Insert(10)
	bst.Insert(5)
	bst.Insert(50)
	bst.Insert(1)
	bst.Insert(40)
	bst.Insert(100)

	bst.Preorder()
	fmt.Println()

	fmt.Println("Number of nodes in range:", getCountOfNode(bst, 5, 45))
}
