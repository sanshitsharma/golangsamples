package main

import (
	"fmt"

	"github.com/sanshitsharma/golangsamples/ds/trees"
)

/*
Print a Binary Tree in Vertical Order | Set 2 (Map based Method)
Given a binary tree, print it vertically. The following example illustrates vertical order traversal.

           1
        /    \
       2      3
      / \   /   \
     4   5  6   7
               /  \
              8   9


The output of print this tree vertically will be:
4
2
1 5 6
3 8
7
9

ref: https://www.geeksforgeeks.org/print-binary-tree-vertical-order-set-2/
*/

var nodesMap = make(map[int][]int)

func verticalOrderRecurse(node *trees.BSTNode, key int) {
	if node == nil {
		return
	}

	verticalOrderRecurse(node.Left, key-1)
	if _, ok := nodesMap[key]; !ok {
		nodesMap[key] = []int{node.Data}
	} else {
		nodesMap[key] = append(nodesMap[key], node.Data)
	}

	verticalOrderRecurse(node.Right, key+1)
}

func verticalOrder(root *trees.BSTNode) {
	verticalOrderRecurse(root, 0)

	for k, v := range nodesMap {
		for _, val := range v {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
		k++
	}
}

func main() {
	bst := trees.NewBST()
	bst = bst.Insert(4)
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(5)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(9)

	fmt.Println("-------PRE-ORDER------")
	bst.Preorder()
	fmt.Println("\n----------------------")

	fmt.Println("\n----- Vertical Order -----")
	verticalOrder(bst)
}
