package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	var low int = 0
	var high int = len(nums) - 1

	for low <= high {

		var mid int = low + ((high - low) / 2)

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		}

		if nums[mid] > target {
			high = mid - 1
		}
	}

	return low
}

func main() {
	fmt.Println(`test`)
	var index int = searchInsert([]int{1, 3, 5, 6}, 2)

	fmt.Println(`index: `, index)
}
