package main

import (
	"github.com/huydinhle/Algorithms-Book-Notes/Data-Structure-And-Algorithms/stack"
)

func matchParentheses(str string) (bool, int) {
	stack := stack.Stack{}

	close := ')'
	open := '('

	for k, v := range str {
		if v == close {
			if stack.IsEmpty() {
				return false, k
			}
			stack.Pop()
		} else if v == open {
			stack.Push(k)
		}
	}
	if stack.IsEmpty() {
		return true, -1
	}
	lastIndex := stack.Pop()
	for !stack.IsEmpty() {
		lastIndex = stack.Pop()
	}
	return false, lastIndex.(int)
}
