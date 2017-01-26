package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSwap(t *testing.T) {
	slice := []int{3, 2, 1}
	Swap(slice, 0, 2)
	if slice[0] != 1 {
		t.Error("Expecting slice[0] is 1 but got", slice[0])
	}
	if slice[2] != 3 {
		t.Error("Expecting slice[2] is 3 but got", slice[2])
	}
}

func TestInsSort(t *testing.T) {
	slice := []int{3, 4, 1, 2}
	InsSort(slice)
	result := sort.IsSorted(sort.IntSlice(slice))
	fmt.Printf("slice = %+v\n", slice)
	if !result {
		t.Error("slice isn't sorted, slice is ", slice)
	}
}

func TestInsSort2(t *testing.T) {
	slice := []int{3, 4, 1, 2}
	InsSort2(slice)
	result := sort.IsSorted(sort.IntSlice(slice))
	fmt.Printf("slice = %+v\n", slice)
	if !result {
		t.Error("slice isn't sorted, slice is ", slice)
	}
}
