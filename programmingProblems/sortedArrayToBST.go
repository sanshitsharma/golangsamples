package main

import "github.com/sanshitsharma/golangsamples/ds/trees"

/*
Given an array which has data members sorted in ascending order. Construct a Balanced Binary Search Tree which
has same data members as the given array.

Example:
Input: Array [1, 2, 3, 4, 5, 6, 7]
Output: A balanced BST
        4
      /   \
     2     6
   /  \   / \
  1   3  5   7

Input: Array [1, 2, 3, 4]
Output: A balanced BST
        2
      /   \
     1     3
   			\
             4
*/

/*
Solution:

At every step, find the middle element of the array which exists at index 'i'.
The mid element is converted into a node of the BST
Then recurse:
low..i-1
i+1..high
*/

func createBSTRecurse(values []int, low, high int) *trees.BSTNode {
	if low > high {
		return nil
	}

	mid := (low + high) / 2

	node := &trees.BSTNode{Data: values[mid], Left: nil, Right: nil}
	node.Left = createBSTRecurse(values, low, mid-1)
	node.Right = createBSTRecurse(values, mid+1, high)

	return node
}

func createBST(values []int) *trees.BSTNode {
	low := 0
	high := len(values) - 1

	return createBSTRecurse(values, low, high)
}

func main() {

	// Input array
	values := []int{1, 2, 3, 4, 5, 6, 7}

	bst := createBST(values)
	bst.Preorder()
}
