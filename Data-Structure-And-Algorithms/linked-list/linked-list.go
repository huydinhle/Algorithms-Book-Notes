package main

import (
	"fmt"
)

// LinkedList type
type LinkedList struct {
	size int
	head *Element
}

// Element is part of linked list
type Element struct {
	value interface{}
	next  *Element
}

// SearchList search for an element in linked list
func (list *LinkedList) SearchList(element interface{}) Element {
	if list.size == 0 {
		return Element{}
	}
	return list.head.SearchElement(element)
}

// SearchElement search for an element in linked list
func (node *Element) SearchElement(element interface{}) Element {
	if node.value == element {
		return *node
	}
	if node.next == nil {
		return Element{}
	}
	return node.next.SearchElement(element)
}

// DeleteNode search for an element in linked list and delete
func (list *LinkedList) DeleteNode(element interface{}) {
	if list.size == 0 {
		return
	}
	result := list.head.SearchPrevElement(element)
	if result.value == nil {
		return
	}
	fmt.Printf("result = %+v\n", result)
	result.next = result.next.next
}

// SearchPrevElement search for a previous element of the specify element
//in linked list
func (node *Element) SearchPrevElement(element interface{}) *Element {
	if node.next.value == element {
		return node
	}
	if node.next == nil {
		return &Element{}
	}
	return node.next.SearchPrevElement(element)
}

// InsertList will insert a item into the list
func (list *LinkedList) InsertList(element interface{}) {
	list.size++
	node := Element{element, nil}
	node.next = list.head
	list.head = &node
}

// PrintList will print the whole linked list out
func (list *LinkedList) PrintList() {
	node := list.head
	fmt.Print("The List: ")
	for node != nil {
		fmt.Print(node.value, " ,")
		node = node.next
	}
	fmt.Println()
}

func main() {
	node := &LinkedList{}
	node.InsertList(4)
	node.InsertList(8)
	node.PrintList()
	node.DeleteNode(4)
	node.PrintList()
}
