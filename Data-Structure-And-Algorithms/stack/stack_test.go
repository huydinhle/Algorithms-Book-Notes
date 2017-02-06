package stack

import (
	"fmt"
	"testing"
)

func TestNothing(t *testing.T) {
	fmt.Println("nothing")
}

func TestStackInsert(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	slice := stack.ToSlice()
	if len(slice) != 3 {
		t.Error("slice's size is not 3, it is ", len(slice))
	}
	if slice[0] != 3 {
		t.Error("slice first position is not 3, it is ", slice[0])
	}
	if slice[1] != 2 {
		t.Error("slice second position is not 2, it is ", slice[1])
	}
	if slice[2] != 1 {
		t.Error("slice third position is not 1, it is ", slice[3])
	}
}

func TestToSliceEmpty(t *testing.T) {
	stack := Stack{}
	if size := stack.list.Size(); size != 0 {
		t.Error("size of stack should be zero, it is ", size)
	}
}

func TestToSliceNotEmpty(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	if size := stack.list.Size(); size != 2 {
		t.Error("size of stack should be 2, it is ", size)
	}
	slice := stack.ToSlice()
	if length := len(slice); length != 2 {
		t.Error("length of stack should be 2, it is ", length)
	}
}

func TestStackPop(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)

	pop := stack.Pop()

	if pop != 2 {
		t.Error("the pop value should be 2 because it is on top of the list. It is ", pop)
	}
	if size := stack.list.Size(); size != 1 {
		t.Error("stack should have the size of . Instead it has the size of ", size)
	}
}
