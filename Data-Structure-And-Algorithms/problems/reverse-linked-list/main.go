package main

import (
	"github.com/huydinhle/Algorithms-Book-Notes/Data-Structure-And-Algorithms/linkedlist"
)

func reverseLinkedList(l linkedlist.LinkedList) linkedlist.LinkedList {
	result := linkedlist.LinkedList{}

	for l.Size() != 0 {
		result.InsertList(l.DeleteNode(l.GetHeadValue()))
	}
	return result
}
