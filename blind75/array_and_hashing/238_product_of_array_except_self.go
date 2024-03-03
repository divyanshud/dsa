package main

import "fmt"

func productExceptSelf(nums []int) []int {
	size := len(nums)
	left := make([]int, size)
	right := make([]int, size)
	left[0] = 1
	right[size-1] = 1
	for idx := 0; idx < size-1; idx++ {
		left[idx+1] = left[idx] * nums[idx]
	}
	for idx := size - 1; idx > 0; idx-- {
		right[idx-1] = right[idx] * nums[idx]
	}

	for idx := range nums {
		nums[idx] = left[idx] * right[idx]
	}
	return nums
}

func main() {
	a := []int{1, 0, 3, 4}
	fmt.Println("Product of elements of array except self is ", productExceptSelf(a))
}
