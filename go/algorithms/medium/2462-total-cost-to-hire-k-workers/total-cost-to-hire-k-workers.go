package main

import (
	"container/heap"
	"fmt"
)

func totalCost(costs []int, k int, candidates int) int64 {
	hired := 0
	sum := 0

	first, last := &MinHeap{}, &MinHeap{}

	for len(*first) < candidates && len(costs) > 0 {
		heap.Push(first, costs[0])
		costs = costs[1:]
	}

	for len(*last) < candidates && len(costs) > 0 {
		heap.Push(last, costs[len(costs)-1])
		costs = costs[:len(costs)-1]
	}

	for hired < k {
		if len(*last) == 0 && len(*first) == 0 {
			return int64(sum)
		}
		hired++

		if len(*last) == 0 && len(*first) != 0 {
			sum += heap.Pop(first).(int)
			continue
		}

		if len(*first) == 0 && len(*last) != 0 {
			sum += heap.Pop(last).(int)
			continue
		}

		lowestFirst := heap.Pop(first).(int)
		lowestLast := heap.Pop(last).(int)

		if lowestFirst <= lowestLast {

			sum += lowestFirst
			heap.Push(last, lowestLast)
			if len(costs) > 0 {
				heap.Push(first, costs[0])
				costs = costs[1:]
			}
		} else {

			sum += lowestLast
			heap.Push(first, lowestFirst)
			if len(costs) > 0 {
				heap.Push(last, costs[len(costs)-1])
				costs = costs[:len(costs)-1]
			}
		}
	}

	return int64(sum)
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Min() interface{} {
	x := h.Pop()
	h.Push(x)
	return x
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	fmt.Println(`test`)
	// test := []int{3, 1, 3, 4, 5, 3}
	result := totalCost([]int{1, 2, 4, 1}, 3, 3)
	// result := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	// result := increasingTriplet([]int{1, 2, 3, 4, 5})
	// result := increasingTriplet([]int{1, 2, 2147483647})

	fmt.Println(result)
}
