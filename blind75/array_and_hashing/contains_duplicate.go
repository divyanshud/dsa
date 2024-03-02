package main

import "fmt"

func containsDuplicate(nums []int) bool {
	freq := make(map[int]bool)
	for _, ele := range nums {
		_, ok := freq[ele]
		if ok {
			return ok
		}
		freq[ele] = !ok
	}
	return false
}

func main() {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(fmt.Sprintf("Array %+v contains duplicate: %t", arr, containsDuplicate(arr)))
}
