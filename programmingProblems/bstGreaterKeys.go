package main

import (
	"fmt"

	"github.com/sanshitsharma/golangsamples/ds/trees"
)

/*
Given a Binary Search Tree (BST), convert it to a Binary Tree such that every key of the original BST
is changed to key plus sum of all greater keys in BST.

Examples:
Input: Root of following BST
              5
            /   \
           2     13

Output: The given BST is converted to following Binary Tree
              18
            /   \
          20     13

*/

/*
This problem can be solved by traversing the BST in recverse inorder and maintaining a sum
field which gets updated after we return from the traverse right call
*/
var sum int

func bstToGreaterKeys(node *trees.BSTNode) {
	if node == nil {
		return
	}

	bstToGreaterKeys(node.Right)
	sum += node.Data
	node.Data = sum
	bstToGreaterKeys(node.Left)
}

func main() {
	// Create a BST first
	bst := trees.NewBST()

	bst = bst.Insert(5)
	bst.Insert(2)
	bst.Insert(13)
	/*
		bst.Insert(4)
		bst.Insert(1)
		bst.Insert(15)
	*/

	bst.Preorder()
	fmt.Println()

	bstToGreaterKeys(bst)

	bst.Preorder()
	fmt.Println()
}
