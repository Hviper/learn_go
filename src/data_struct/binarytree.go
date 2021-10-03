package main

import "fmt"

// BTree Returns a binary tree structure which contains only a root Node
type BTree struct {
	Root *Node
}

// Insert a value in the BTree
func Insert(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.left = Insert(root.left, val)
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}

func VisitNode(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	VisitNode(root.left)
	VisitNode(root.right)
}

func main() {
	root := NewNode(15)
	root.left = NewNode(14)
	root.right = NewNode(13)
	VisitNode(root)
}
