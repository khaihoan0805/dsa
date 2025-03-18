package main

import "container/heap"

type SmallestInfiniteSet struct {
	smallest int
	added    *IntHeap
	inHeap   map[int]bool // the value : whether the value is inside the heap
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{0, h, map[int]bool{}}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.added.Len() > 0 {
		pop := heap.Pop(this.added).(int)
		this.inHeap[pop] = false
		return pop
	}

	this.smallest++
	return this.smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num <= this.smallest && !this.inHeap[num] {
		this.inHeap[num] = true
		heap.Push(this.added, num)
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// fmt.Println(`test`)
	// result := findKthLargest([]int{2, 2, 2, 1, 1, 1, 5}, 2)

	// fmt.Println(result)

	smallestInfiniteSet := Constructor()
	smallestInfiniteSet.AddBack(2)    // 2 is already in the set, so no change is made.
	smallestInfiniteSet.PopSmallest() // return 1, since 1 is the smallest number, and remove it from the set.
	smallestInfiniteSet.PopSmallest() // return 2, and remove it from the set.
	smallestInfiniteSet.PopSmallest() // return 3, and remove it from the set.
	smallestInfiniteSet.AddBack(1)    // 1 is added back to the set.
	smallestInfiniteSet.PopSmallest() // return 1, since 1 was added back to the set and
	// is the smallest number, and remove it from the set.
	smallestInfiniteSet.PopSmallest() // return 4, and remove it from the set.
	smallestInfiniteSet.PopSmallest() // return 5, and remove it from the set.
}
