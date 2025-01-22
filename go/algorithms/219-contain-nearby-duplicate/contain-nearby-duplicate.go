package main

import (
	"fmt"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	reminder := make(map[int]int)

	for i, v := range nums {
		value, ok := reminder[v]

		if ok && int(math.Abs(float64(i-value))) <= k {
			return true
		} else {
			reminder[v] = i
		}
	}

	return false
}

func main() {
	fmt.Println(`test`)

	result := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)

	fmt.Println(`result: `, result)
}
