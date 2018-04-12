package main

import (
	"fmt"
	"math"

	"github.com/sanshitsharma/golangsamples/ds/trees"
)

func getTreeHeight(bst *trees.BSTNode) float64 {
	if bst == nil {
		return 0.00
	}

	return 1.00 + math.Max(getTreeHeight(bst.Left), getTreeHeight(bst.Right))
}

func main() {
	bst := trees.NewBST()

	bst = bst.Insert(20)
	bst.Insert(8)
	bst.Insert(22)
	bst.Insert(4)
	bst.Insert(12)
	bst.Insert(10)
	bst.Insert(14)
	bst.Insert(9)

	height := getTreeHeight(bst)
	fmt.Println("---- Preorder Traversal ----")
	bst.Preorder()
	fmt.Println("\n----------------------------")

	fmt.Println("Height of Tree:", height)
}
