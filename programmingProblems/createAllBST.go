package main

import (
	"github.com/sanshitsharma/golangsamples/ds/trees"
)

/*
Construct all possible BSTs for keys 1 to N
*/

func contructTrees(low, high int) *trees.BSTNode {
	if low > high {
		return nil
	}

	node := &trees.BSTNode{Data: low, Left: nil, Right: nil}
	node.Left = contructTrees(1, low - 1)
	node.Right = contructTrees(low+1, high)

	return node
}

func constructAllTreesUtil(bstArr []*trees.BSTNode, low, high int) {
	for i := low; i <= high; i++ {
		contructTrees(low, high)
	}
}

func constructAllTrees(n int) []*trees.BSTNode {
	nBst := make([]*trees.BSTNode, 0)
	low := 1
	high := n

	constructAllTreesUtil(nBst, low, high)
}

func main() {
	n := 3
	nBsts := constructAllTrees(n)
}