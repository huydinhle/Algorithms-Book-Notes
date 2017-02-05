package main

import (
	"fmt"
	"testing"
)

func testInsert(t *testing.T) {
	node := generateTestList()
	fmt.Printf("node = %+v\n", node)
}

func generateTestList() LinkedList {
	node := LinkedList{1, nil}
	node.InsertList(2)
	node.InsertList(3)
	node.InsertList(4)

}
