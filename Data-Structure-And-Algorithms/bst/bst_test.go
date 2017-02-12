package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := Bst{}
	tree.Insert(7, 2)
	tree.Insert(3, 3)
	tree.Insert(1, 1)
	tree.Insert(2, 1)
	tree.InOrderRec()
	fmt.Println()
}
