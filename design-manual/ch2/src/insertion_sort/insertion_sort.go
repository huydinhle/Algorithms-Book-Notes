package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// InsSort is insertion sort
func InsSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if nums[j] < nums[j-1] {
				Swap(nums, j, j-1)
			}
			fmt.Printf("nums = %+v\n", nums)
		}
	}
}

//InsSort2 is insertion sort but we swap a
// bit backward comparing to the above methods
// the j instead going from i+1 to zero
// it will go from i+1 to len(nums)
func InsSort2(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				Swap(nums, i, j)
			}
		}
	}

}

// Swap swaps two elements inside a slice
func Swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
