package main

import (
	"fmt"

	"github.com/sanshitsharma/golangsamples/ds/queue"
	"github.com/sanshitsharma/golangsamples/ds/trees"
)

// LevelOrder returns an array containing the BST elements traversed level by level
func LevelOrder(bst *trees.BSTNode) []int {
	// We will use a queue to process the tree nodes
	// This is similar to Breadth first traversal

	res := make([]int, 0)
	q := queue.NewLinkQueue()
	q.Enqueue(bst)

	for !q.IsEmpty() {
		node, _ := q.Dequeue().(*trees.BSTNode)
		res = append(res, node.Data)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}

		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}

	return res
}

// LevelOrderPrint prints the elements of a binary tree level by level
// with each level printed on a new line
func LevelOrderPrint(bst *trees.BSTNode) {
	// The idea is to essentially use two queues, oddQ & evenQ and enqueue
	// & dequeue from one of the queues at each level
	oddQ := queue.NewLinkQueue()
	evenQ := queue.NewLinkQueue()

	// Typically root level is known as level0 so we will treat it an
	// even level and push to evenQ
	evenQ.Enqueue(bst)

	for !(oddQ.IsEmpty() && evenQ.IsEmpty()) {
		// First read from evenQ.
		for !evenQ.IsEmpty() {
			node, _ := evenQ.Dequeue().(*trees.BSTNode)
			fmt.Printf("%v ", node.Data)
			if node.Left != nil {
				oddQ.Enqueue(node.Left)
			}
			if node.Right != nil {
				oddQ.Enqueue(node.Right)
			}
		}
		fmt.Println()

		for !oddQ.IsEmpty() {
			node, _ := oddQ.Dequeue().(*trees.BSTNode)
			fmt.Printf("%v ", node.Data)

			if node.Left != nil {
				evenQ.Enqueue(node.Left)
			}
			if node.Right != nil {
				evenQ.Enqueue(node.Right)
			}
		}
		fmt.Println()
	}
}

func main() {
	bst := trees.NewBST()
	bst = bst.Insert(7)
	bst.Insert(4)
	bst.Insert(12)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(8)
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(10)

	//bst.Preorder()
	//fmt.Println()

	trav := LevelOrder(bst)
	fmt.Println(trav)

	fmt.Printf("\nPrint level by level\n")
	LevelOrderPrint(bst)
}
