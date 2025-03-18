package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	arr := [][]int{}
	for i := 0; i < len(nums1); i++ {
		arr = append(arr, []int{nums1[i], nums2[i]})
	}

	// O(NlogN)
	// Sort in decreasing order of nums2
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	minHeap := &MinHeap{}
	sum := 0
	// O(KlogK)
	for i := 0; i < k; i++ {
		// O(logK)
		heap.Push(minHeap, arr[i])
		sum += arr[i][0]
	}

	res := sum * arr[k-1][1]

	// O(NlogK)
	for i := k; i < len(arr); i++ {
		// Remove element with smallest nums1
		// O(logK)
		smallNum := heap.Pop(minHeap).([]int)
		sum += arr[i][0] - smallNum[0]

		// O(logK)
		heap.Push(minHeap, arr[i])

		res = max(res, sum*arr[i][1])
	}

	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	fmt.Println(`test`)
	// test := []int{3, 1, 3, 4, 5, 3}
	result := maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 2)
	// result := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	// result := increasingTriplet([]int{1, 2, 3, 4, 5})
	// result := increasingTriplet([]int{1, 2, 2147483647})

	fmt.Println(result)
}
