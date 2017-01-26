package main

import (
	"sort"
	"testing"
)

func TestSwap(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1, 0, 10}
	Swap(slice, 0, 6)
	if v := slice[0]; v != 10 {
		t.Error("exepecting 10 for the first position but got", v)
	}
	if v := slice[6]; v != 5 {
		t.Error("exepecting 5 for the 6h position but got", v)
	}

}

func TestSelectionsort(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1, 0, 10}
	Selectionsort(slice)
	isSorted := sort.IsSorted(sort.IntSlice(slice))
	// fmt.Printf("slice = %+v\n", slice)
	// fmt.Printf("isSorted = %+v\n", isSorted)
	if !isSorted {
		t.Error("Expecting the Array to be sorted but it is", slice)
	}
}
