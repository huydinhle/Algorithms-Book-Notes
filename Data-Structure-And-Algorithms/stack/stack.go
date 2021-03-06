package stack

import (
	"github.com/huydinhle/Algorithms-Book-Notes/Data-Structure-And-Algorithms/linkedlist"
)

// Stack implement LIFO data structure
type Stack struct {
	list linkedlist.LinkedList
}

// ToSlice convert the current stack into
// a slice. This is mainly for testing
func (stack *Stack) ToSlice() []interface{} {
	if stack.list.Size() == 0 {
		return []interface{}{}
	}
	return stack.list.ToSlice()
}

// Push will add an item on top the stack
func (stack *Stack) Push(element interface{}) {
	stack.list.InsertList(element)
}

// Pop will pop the item from the top of the list
func (stack *Stack) Pop() interface{} {
	poppedNode := stack.list.DeleteNode(stack.list.GetHeadValue())
	return poppedNode
}

// IsEmpty let you know if the stack is empty or not
func (stack *Stack) IsEmpty() bool {
	if stack.list.Size() == 0 {
		return true
	}
	return false
}
