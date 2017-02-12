package main

import (
	"fmt"
)

// TreeNode is a representation of a node in a tree
// it has left node and right node and a value
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	key   int
	value int
}

// Bst Binary  Search Tree Implementation
// Please write implementation
// At the end please put time complexity for each
// operations
type Bst struct {
	head *TreeNode
}

// Insert into the BST
func (t *Bst) Insert(key, value int) {
	node := t.head
	if t.head == nil {
		t.head = &TreeNode{nil, nil, key, value}
		return
	}
	for {
		if key > node.key {
			if node.right == nil {
				node.right = &TreeNode{nil, nil, key, value}
				return
			}
			node = node.right
		} else {
			if node.left == nil {
				node.left = &TreeNode{nil, nil, key, value}
				return
			}
			node = node.left
		}
	}
}

// Search will seach for an item with key
// and return the value
func (t *Bst) Search(key int) int {

	return 0
}

// Delete will seach for an item with key
// delete the item, and then return the value
func (t *Bst) Delete(key int) int {
	return 0
}

// FindMax returns the value of the max key
func (t *Bst) FindMax() int {
	return 0
}

// InOrderRec in-order recursively traversal
func (t *Bst) InOrderRec() {
	inorderRecHelper(t.head)
}

func inorderRecHelper(node *TreeNode) {
	if node == nil {
		return
	}
	inorderRecHelper(node.left)
	fmt.Print(node.value, " ")
	inorderRecHelper(node.right)
}

// PreOrderRec pre-order recursively traversal
func (t *Bst) PreOrderRec() {
}

// PostOrderRec post-order recursively traversal
func (t *Bst) PostOrderRec() {
}

// InOrderIter in-order iteratively == DFS1
func (t *Bst) InOrderIter() {
}

// PreOrderIter pre-order iteratively == DFS2
func (t *Bst) PreOrderIter() {
}

// PostOrderIter pre-order iteratively == DFS3
func (t *Bst) PostOrderIter() {
}

// BFS Breath First Search
func (t *Bst) BFS() {
}
