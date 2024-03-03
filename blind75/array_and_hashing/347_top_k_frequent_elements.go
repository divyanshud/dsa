package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
	We will solve this using 2 solutions.
	1. Making a freq map and then sorting it in decreasing order.
	   Iterate over map k times to fetch top k frequent elements. Complexity is O(nLog(n)) because of sorting
	2. Making a Max Heap based on freq of elements and then popping out root node k times. Complexity is O(nLog(n))
	   because of building heap.
*/

// 1st Solution

type Element struct {
	key   int
	count int
}

type ByValue []Element

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].count > a[j].count }

func findTopKFreqElementsUsingSort(arr []int, k int) []int {
	freq := make(map[int]int)
	for _, ele := range arr {
		freq[ele] += 1
	}

	// Create a slice of key-value pairs.
	var kv []Element
	for k, v := range freq {
		kv = append(kv, Element{k, v})
	}

	// Sort the slice by value.
	sort.Sort(ByValue(kv))
	results := make([]int, k)
	for i := 0; i < k; i++ {
		results[i] = kv[i].key
	}
	return results
}

/*
	2nd Solution
*/

type MaxHeap []Element

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].count > m[j].count
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(Element))
}

func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

// Using Max heap
func findTopKFreqElementsUsingHeap(arr []int, k int) []int {
	freq := make(map[int]int)
	for _, ele := range arr {
		freq[ele] += 1
	}

	mh := MaxHeap{}

	for k, v := range freq {
		mh.Push(Element{
			key:   k,
			count: v,
		})
	}

	// Call this whenever heap in invalidated which will internally do heapify
	heap.Init(&mh)

	results := make([]int, k)
	for i := 0; i < k; i++ {
		results[i] = heap.Pop(&mh).(Element).key
	}
	return results
}

func main() {
	arr := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(fmt.Sprintf("Top k=%d frequent elements using sorting a map in array %+v is: %+v", k, arr, findTopKFreqElementsUsingSort(arr, k)))
	fmt.Println(fmt.Sprintf("Top k=%d frequent elements using MaxHeap in array %+v is: %+v", k, arr, findTopKFreqElementsUsingHeap(arr, k)))
}
