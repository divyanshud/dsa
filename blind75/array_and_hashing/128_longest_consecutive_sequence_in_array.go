package main

import (
	"fmt"
	"math"
)

func longestConsecutive(nums []int) int {
	freq := make(map[int]bool)
	for idx := range nums {
		freq[nums[idx]] = true
	}

	length := math.MinInt64
	for key := range freq {
		temp := 1
		if freq[key-1] == false {
			for freq[key+temp] == true {
				temp++
			}
		}
		length = max(length, temp)
	}
	return length
}

func main() {
	arr := []int{100, 4, 200, 1, 3, 2, 7, 11, 9, 8, 12, 14, 13, 10}
	fmt.Println("The longest consecutive seq in array is:", longestConsecutive(arr))
}
