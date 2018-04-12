package main

import (
	"fmt"
	"math"

	"github.com/sanshitsharma/golangsamples/ds/trees"
)

func createMirror(node *trees.BSTNode) {
	if node == nil {
		return
	}

	temp := node.Left
	node.Left = node.Right
	node.Right = temp

	createMirror(node.Left)
	createMirror(node.Right)
}

func getHeight(bst *trees.BSTNode) int {
	if bst == nil {
		return 0
	}

	return 1 + int(math.Max(float64(getHeight(bst.Left)), float64(getHeight(bst.Right))))
}

func isBalanced(bst *trees.BSTNode) bool {
	if bst == nil {
		return true
	}

	leftHeight := getHeight(bst.Left)
	rightHeight := getHeight(bst.Right)

	fmt.Printf("Left Height = '%v', Right Height = '%v'\n", leftHeight, rightHeight)
	if int(math.Abs(float64(leftHeight-rightHeight))) > 1 {
		return false
	}

	return true
}

func isBalancedOptimized(root *trees.BSTNode, height *int) bool {
	if root == nil {
		return true
	}

	lh := 0 // Height of left subtree
	rh := 0 // Height of right subtree

	l := isBalancedOptimized(root.Left, &lh)  // Will be true is left subtree is balanced
	r := isBalancedOptimized(root.Right, &rh) // Will be true if right subtree is balanced

	if lh >= rh {
		*height = lh + 1
	} else {
		*height = rh + 1
	}

	if int(math.Abs(float64(lh-rh))) > 1 {
		return false
	}

	return l && r
}

func main() {
	bst := trees.NewBST()
	bst = bst.Insert(4)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(8)

	fmt.Println("-------- Preorder --------")
	bst.Preorder()
	fmt.Println("\n-------------------------")

	createMirror(bst)

	fmt.Println("-------- Preorder --------")
	bst.Preorder()
	fmt.Println("\n-------------------------")

	//fmt.Println("Is Balanced:", isBalanced(bst))
	height := 0
	fmt.Println("Is Balanced:", isBalancedOptimized(bst, &height))
}
