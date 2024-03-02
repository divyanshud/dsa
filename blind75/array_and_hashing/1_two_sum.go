package main

import "fmt"

func twoSum(nums []int, target int) []int {
	freq := make(map[int]int)
	for idx1, ele := range nums {
		idx2, ok := freq[target-ele]
		if ok {
			return []int{idx1, idx2}
		}
		freq[ele] = idx1
	}
	return []int{}
}

func main() {
	arr := []int{2, 7, 11, 15}
	sum := 9
	fmt.Println(fmt.Sprintf("The element index with sum %d are: %+v", sum, twoSum(arr, sum)))
}
