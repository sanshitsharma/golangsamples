package trees

import (
	"errors"
	"fmt"

	"github.com/sanshitsharma/golangsamples/ds/stack"
)

var (
	errInvalidParams = errors.New("invalid params")
	errBadDataType   = errors.New("unsupported data type for BST")
	errKeyNotFound   = errors.New("key not found in tree")
)

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// BSTNode is the node structure for binary search tree
type BSTNode struct {
	Data  int
	Left  *BSTNode
	Right *BSTNode
}

// NewBST creates a new binary search tree
func NewBST() *BSTNode {
	return nil
}

// Insert adds a new value to BST
func (bst *BSTNode) Insert(value int) *BSTNode {
	if bst == nil {
		return &BSTNode{
			Data:  value,
			Left:  nil,
			Right: nil,
		}
	}

	if bst.Data == minInt {
		bst.Data = value
		return nil
	}

	if value < bst.Data {
		bst.Left = bst.Left.Insert(value)
	} else {
		bst.Right = bst.Right.Insert(value)
	}

	return bst
}

// Search looks for the key in tree. If found, it returns the node
func (bst *BSTNode) Search(value int) (*BSTNode, error) {
	if bst == nil {
		return nil, errKeyNotFound
	}

	if bst.Data == value {
		return bst, nil
	} else if bst.Data > value {
		return bst.Left.Search(value)
	} else {
		return bst.Right.Search(value)
	}
}

func findParentRecurse(node *BSTNode, parent *BSTNode, value int) *BSTNode {
	if value == node.Data {
		return parent
	}

	if value < node.Data {
		return findParentRecurse(node.Left, node, value)
	}
	return findParentRecurse(node.Right, node, value)
}

// Parent returns the parent of the node is it exists in the
// BST
func (bst *BSTNode) Parent(value int) (*BSTNode, error) {
	if _, err := bst.Search(value); err != nil {
		return nil, errKeyNotFound
	}
	var parent *BSTNode
	return findParentRecurse(bst, parent, value), nil
}

func inorderSuccessorOfDeleteNode(node *BSTNode) *BSTNode {
	curr := node.Right
	for curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func (bst *BSTNode) deleteRecurse(node, parent *BSTNode) *BSTNode {
	// if node was found, delete it
	if node.Left == nil && node.Right == nil {
		// It's a leaf node. This can be directly deleted.
		if parent == nil {
			// node is the root node
			return parent
		}

		if parent.Left == node {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if node.Left != nil && node.Right != nil {
		// Node has two children. We need to find the inorder successor
		iSucc := inorderSuccessorOfDeleteNode(node)
		p, _ := bst.Parent(iSucc.Data)
		node.Data = iSucc.Data
		return bst.deleteRecurse(iSucc, p)
	} else {
		// Node has just one child. Replace node value with child and delete child
		if node.Left != nil {
			node.Data = node.Left.Data
			return bst.deleteRecurse(node.Left, node)
		}
		node.Data = node.Right.Data
		return bst.deleteRecurse(node.Right, node)
	}

	return bst
}

// Delete removes the node from BST if it exists
func (bst *BSTNode) Delete(value int) (*BSTNode, error) {
	node, err := bst.Search(value)
	if err != nil {
		return nil, err
	}

	// Find the parent for node
	parent, _ := bst.Parent(value)
	bst = bst.deleteRecurse(node, parent)
	return bst, nil
}

// Inorder prints the in-order traversal of the BST
func (bst *BSTNode) Inorder() {
	if bst == nil {
		return
	}

	bst.Left.Inorder()
	fmt.Printf("%v ", bst.Data)
	bst.Right.Inorder()
}

// Preorder prints the pre order traversal of the tree
func (bst *BSTNode) Preorder() {
	if bst == nil {
		return
	}

	fmt.Printf("%v ", bst.Data)
	bst.Left.Preorder()
	bst.Right.Preorder()
}

func contructTreeUtil(node *BSTNode, values []int, low, high int) {

	//fmt.Println("Contructing for", node.Data, "Low =", low, "High =", high)

	// Termination condition
	if low > high {
		return
	}

	// Find the first element that is greater than node.Data
	i := low
	for i < high && values[i] < node.Data {
		i++
	}

	//fmt.Println("Found next greater element at index:", i) //, "Value:", values[i])

	// Set left child if the next element is not i
	if low != i {
		node.Left = &BSTNode{Data: values[low], Left: nil, Right: nil}
		// Go left
		contructTreeUtil(node.Left, values, low+1, i-1)
	}

	if low <= high {
		// Set right child
		node.Right = &BSTNode{Data: values[i], Left: nil, Right: nil}
		// Go right
		contructTreeUtil(node.Right, values, i+1, high)
	}
}

// ConstructTree creates a BST from a given preorder traversal
func ConstructTree(values []int) *BSTNode {
	if len(values) == 0 {
		return nil
	}

	// Contruct the root node of the tree from the first element of the array
	root := &BSTNode{
		Data:  values[0],
		Left:  nil,
		Right: nil,
	}

	contructTreeUtil(root, values, 1, len(values)-1)

	return root
}

// Helper func to fetch value from stack
func getValueFromStack(stk *stack.Stack, opType string) *BSTNode {
	var node *BSTNode

	switch opType {
	case "peek":
		elem, err := stk.Peek()
		if err != nil {
			fmt.Println("failed to peek into stack. err:", err)
			return nil
		}

		var ok bool
		node, ok = elem.(*BSTNode)
		if !ok {
			fmt.Println("peeked element is not of type *BSTNode")
		}
	case "pop":
		elem, err := stk.Pop()
		if err != nil {
			fmt.Println("failed to peek into stack. err:", err)
			return nil
		}

		var ok bool
		node, ok = elem.(*BSTNode)
		if !ok {
			fmt.Println("peeked element is not of type *BSTNode")
		}
	default:
		fmt.Println("unsuporrted opType:", opType)
	}

	return node
}

// ConstructTreeIter creates a BST from a given preorder traversal in
// O(n) time and uses a stack
func ConstructTreeIter(values []int) *BSTNode {
	if values == nil {
		return nil
	}

	// Create the root node
	root := &BSTNode{Data: values[0], Left: nil, Right: nil}

	// Create a stack & push root node to it
	stk := stack.NewStack()
	stk.Push(root)

	for i := 1; i < len(values); i++ {
		// If the value is smaller than top of stack, then make the new
		// value a left child of the top and push the new value into stack
		node := getValueFromStack(stk, "peek")
		if node == nil {
			fmt.Println("Stack is empty..")
		}

		if node.Data > values[i] {
			node.Left = &BSTNode{Data: values[i], Left: nil, Right: nil}
			stk.Push(node.Left)
			continue
		}

		// Othewise pop the stack while the new value is > top of stack.
		// Then add the new value as the right child of the last popped value
		// & push the new value onto the stack
		lastNode := getValueFromStack(stk, "peek")

		for !stk.IsEmpty() && node.Data < values[i] {
			lastNode = getValueFromStack(stk, "pop")
			node = getValueFromStack(stk, "peek")
		}

		// Make new value the right child of last popped node and the add the new node
		// to the stack
		lastNode.Right = &BSTNode{Data: values[i], Left: nil, Right: nil}
		stk.Push(lastNode.Right)
	}

	return root
}