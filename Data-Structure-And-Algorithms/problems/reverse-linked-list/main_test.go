package main

import (
	"fmt"
	"github.com/huydinhle/Algorithms-Book-Notes/Data-Structure-And-Algorithms/linkedlist"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	input := linkedlist.LinkedList{}
	input.InsertListEnd(4)
	input.InsertListEnd(3)
	input.InsertListEnd(2)
	input.InsertListEnd(1)
	reverseList := reverseLinkedList(input)
	result := reverseList.ToSlice()
	fmt.Printf("result = %+v\n", result)
}
