package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	operations := 0
	left, right := 0, len(nums)-1

	for right > left {
		sum := nums[left] + nums[right]

		if sum == k {
			operations++
			left++
			right--
		} else if sum > k {
			right--
		} else {
			left++
		}
	}

	return operations
}

func main() {
	fmt.Println(`test`)
	test := []int{3, 1, 3, 4, 5, 3}
	result := maxOperations(test, 6)
	// result := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	// result := increasingTriplet([]int{1, 2, 3, 4, 5})
	// result := increasingTriplet([]int{1, 2, 2147483647})

	fmt.Println(result)
}
