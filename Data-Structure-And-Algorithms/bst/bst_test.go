package main

import (
	// "fmt"
	"reflect"
	"sort"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := Bst{}
	tree.Insert(7, 2)
	tree.Insert(3, 3)
	tree.Insert(1, 1)
	tree.Insert(2, 1)
	slice := []int{}
	f := func(key int) {
		slice = append(slice, key)
	}
	tree.InOrderRec(f)
	check := sort.IsSorted(sort.IntSlice(slice))
	if !check {
		t.Error("the slice is not sorted, it is ", slice)
	}
}

func TestInsertAlotInOrder(t *testing.T) {
	tree := Bst{}
	tree.Insert(5, 1)
	tree.Insert(2, 2)
	tree.Insert(7, 3)
	tree.Insert(1, 4)
	tree.Insert(4, 5)
	tree.Insert(6, 6)
	tree.Insert(9, 7)
	tree.Insert(8, 8)
	tree.Insert(10, 9)
	slice := []int{}
	f := func(key int) {
		slice = append(slice, key)
	}
	tree.InOrderRec(f)
	check := sort.IsSorted(sort.IntSlice(slice))
	if !check {
		t.Error("the slice is not sorted, it is ", slice)
	}
}

func TestInsertAlotPreOrder(t *testing.T) {
	tree := Bst{}
	tree.Insert(5, 1)
	tree.Insert(2, 2)
	tree.Insert(7, 3)
	tree.Insert(1, 4)
	tree.Insert(4, 5)
	tree.Insert(6, 6)
	tree.Insert(9, 7)
	tree.Insert(8, 8)
	tree.Insert(10, 9)
	slice := []int{}
	f := func(key int) {
		slice = append(slice, key)
	}
	tree.PreOrderRec(f)
	reflect.DeepEqual(slice, []int{5, 2, 1, 4, 7, 6, 9, 8, 10})
}

func TestInsertAlotPostOrder(t *testing.T) {
	tree := Bst{}
	tree.Insert(5, 1)
	tree.Insert(2, 2)
	tree.Insert(7, 3)
	tree.Insert(1, 4)
	tree.Insert(4, 5)
	tree.Insert(6, 6)
	tree.Insert(9, 7)
	tree.Insert(8, 8)
	tree.Insert(10, 9)
	slice := []int{}
	f := func(key int) {
		slice = append(slice, key)
	}
	tree.PostOrderRec(f)
	reflect.DeepEqual(slice, []int{1, 4, 2, 6, 8, 10, 9, 7, 5})
}

func TestSearchFound(t *testing.T) {
	tree := Bst{}
	tree.Insert(5, 1)
	tree.Insert(2, 2)
	tree.Insert(7, 3)
	tree.Insert(1, 4)
	tree.Insert(4, 5)
	tree.Insert(6, 6)
	tree.Insert(9, 7)
	tree.Insert(8, 8)
	tree.Insert(10, 9)
	found, value := tree.Search(10)

	if !found {
		t.Error("key not found, it should be found")
	}
	if value != 9 {
		t.Error("value should be 9 but it is ", value)
	}
}

func TestSearchNotFound(t *testing.T) {
	tree := Bst{}
	tree.Insert(5, 1)
	tree.Insert(2, 2)
	tree.Insert(7, 3)
	tree.Insert(1, 4)
	tree.Insert(4, 5)
	tree.Insert(6, 6)
	tree.Insert(9, 7)
	tree.Insert(8, 8)
	tree.Insert(10, 9)
	found, _ := tree.Search(0)

	if found {
		t.Error("key should not be found but it is found.")
	}
}
