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
