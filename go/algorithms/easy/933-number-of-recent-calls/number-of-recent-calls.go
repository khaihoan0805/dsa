package main

import (
	"fmt"
)

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)

	smallest := -1

	for i := 0; i < len(this.queue); i++ {
		if this.queue[i] > t-3000 {
			smallest = i
			break
		}
	}

	return len(this.queue[smallest:])
}

func main() {
	fmt.Println(`test`)
	test := []int{642, 1849, 4921, 5936, 5957}
	var result []int

	recentCounter := RecentCounter{}
	for _, v := range test {
		sub := recentCounter.Ping(v)
		result = append(result, sub)
	}

	fmt.Println(result)
}
