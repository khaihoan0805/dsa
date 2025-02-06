package main

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	deleted := false
	deletedIndex := -1

	max := 0
	var slice []int = nums[:0]
	// count := 0

	for index := 0; index < len(nums); index++ {
		if nums[index] == 1 {
			slice = append(slice, 1)
			// count++
		}

		if nums[index] == 0 {
			if !deleted {
				deleted = true
				deletedIndex = index
			} else {
				slice = nums[deletedIndex+1 : index]
				// count = index - deletedIndex
				deletedIndex = index
			}
		}

		if len(slice) > max {
			max = len(slice)
		}
	}

	if !deleted {
		max--
		// count--
	}

	// fmt.Println(count)

	return max
}

func main() {
	fmt.Println(`test`)
	// result := longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	// result := longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	result := longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})

	fmt.Println(result)
	// fmt.Println(result1)
}
