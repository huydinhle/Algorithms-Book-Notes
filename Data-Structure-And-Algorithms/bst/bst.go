package main

import (
	// "fmt"
	"reflect"
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
	// O(logn)
}

// Search will seach for an item with key
// and return the value
func (t *Bst) Search(key int) (bool, int) {
	if t.head == nil {
		return false, 0
	}

	node := t.search(key)

	if node == nil {
		return false, 0
	}
	return true, node.value
	// O(logn)
}

func (t *Bst) search(key int) *TreeNode {
	node := t.head
	for node != nil && node.key != key {
		if key > node.key {
			node = node.right
		} else {
			node = node.left
		}
	}
	return node
}

func (t *Bst) searchPrev(key int) *TreeNode {
	node := t.head
	var prev *TreeNode
	prev = nil
	for node != nil && node.key != key {
		if key > node.key {
			prev = node
			node = node.right
		} else {
			prev = node
			node = node.left
		}
	}
	return prev
}

// Delete will seach for an item with key
// delete the item, and then return the value
func (t *Bst) Delete(key int) (bool, int) {
	prev := t.searchPrev(key)
	found := t.search(key)
	if found == nil {
		return false, 0
	}

	///////////////////////
	// no child on found //
	///////////////////////

	if found.left == nil && found.right == nil {
		if prev == nil { // prev == nil means we are trying to delete the head node
			t.head == nil
			return true, found.value
		}
		// found is not head
		if prev.left.key == found.key { // found on the left of previous
			prev.left = nil
		} else { // found on the right of previous
			prev.right = nil
		}
		return true, found.value
	}

	///////////////////////////////
	// exactly one child on found//
	///////////////////////////////

	if found.left == nil || found.right == nil {
		// found is the t.head
		if prev == nil { // meaning this is the head node
			if found.left == nil {
				t.head = found.right
			} else {
				t.head = foudn.left
			}
			return true, found.value
		}
		// found is not head
		if prev.left.key == found.key { // found on the left of previous
			if found.left == nil { // left of found is nil,
				prev.left = found.right
			} else {
				prev.left = found.left
			}
		} else { // found on the right of previous
			if found.left == nil { // left of found is nil,
				prev.right = found.right
			} else {
				prev.right = found.left
			}
		}
		return true, found.value
	}

	///////////////////////////////////
	// two child on found /////////////
	///////////////////////////////////

	max, previous := findMaxUnderNode(found.left)
	found.value = max.value
	found.key = max.key
	if previous != nil {
		previous.right = max.left
	} else {
		found.left = nil
	}
	return true, found.value
}

func (t *Bst) findMaxUnderNode(node *TreeNode) (*TreeNode, *TreeNode) {
	var prev *TreeNode
	prev = nil
	for node.right != nil {
		previous = node
		node = node.right
	}
	return node, prev
}

// FindMax returns the value of the max key
func (t *Bst) FindMax() int {
	return 0
}

// InOrderRec in-order recursively traversal
func (t *Bst) InOrderRec(f func(key int)) {
	inorderRecHelper(t.head, f)
	// O(n)
}

func inorderRecHelper(node *TreeNode, f func(key int)) {
	if node == nil {
		return
	}
	inorderRecHelper(node.left, f)
	f(node.key)
	inorderRecHelper(node.right, f)
}

// PreOrderRec pre-order recursively traversal
func (t *Bst) PreOrderRec(f func(key int)) {
	preorderRecHelper(t.head, f)
	// This method will take the O(n)
}

func preorderRecHelper(node *TreeNode, f func(key int)) {
	if node == nil {
		return
	}
	f(node.key)
	preorderRecHelper(node.left, f)
	preorderRecHelper(node.right, f)
}

// PostOrderRec post-order recursively traversal
func (t *Bst) PostOrderRec(f func(key int)) {
	postorderRecHelper(t.head, f)
	// This method will take the O(n)
}

func postorderRecHelper(node *TreeNode, f func(key int)) {
	if node == nil {
		return
	}
	postorderRecHelper(node.left, f)
	postorderRecHelper(node.right, f)
	f(node.key)
}

// InOrderIter in-order iteratively == DFS1 Stack
func (t *Bst) InOrderIter() {
}

// PreOrderIter pre-order iteratively == DFS2 Stack
func (t *Bst) PreOrderIter() {
}

// PostOrderIter pre-order iteratively == DFS3 Stack
func (t *Bst) PostOrderIter() {
}

// BFS Breath First Search using a Queue
func (t *Bst) BFS() {
}
