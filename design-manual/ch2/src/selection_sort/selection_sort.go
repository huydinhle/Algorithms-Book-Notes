package main

import (
	"fmt"
)

//Selectionsort is the function that will sort
// a slice of numer in increasing order
// The algorithm picks the smallest numer
// each time it iterates through the whole
// slice and put the smallest to the front
// of the slice
func Selectionsort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i + 1
		for j := i; j < len(nums); j++ {
			// fmt.Printf("i = %+v\n", i)
			if nums[j] < min {
				min = j
			}
		}
		// fmt.Printf("i = %+v\n", i)
		if nums[i] > nums[min] {
			// fmt.Printf("inner i = %+v\n", i)
			// fmt.Printf("min = %+v\n", min)
			Swap(nums, i, min)
		}
		fmt.Printf("nums = %+v\n", nums)
	}
}

// Swap will swap two elements inside a slice
func Swap(nums []int, index1 int, index2 int) {
	nums[index1], nums[index2] = nums[index2], nums[index1]
}
