package main

import (
// "fmt"
)

// TreeNode is a representation of a node in a tree
// it has left node and right node and a value
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

// Bst Binary  Search Tree Implementation
type Bst struct {
	head *TreeNode
}

// Insert into the BST
func (t *Bst) Insert(value int) {
}

// TraverseAndPrintPretty traverse in-order and print the tree out
// pretty
func (t *Bst) TraverseAndPrintPretty() {
}
