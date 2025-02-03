package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}

	start := 0
	end := 0
	for index := 1; index < len(nums); index++ {
		if nums[index] != nums[end]+1 {
			if end == start {
				result = append(result, strconv.Itoa(nums[start]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end]))
			}

			start = index
			end = index
			continue
		}

		end = index
	}

	if end == start {
		result = append(result, strconv.Itoa(nums[start]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end]))
	}

	return result
}

func main() {
	fmt.Println(`test`)
	// Define a map
	result := summaryRanges([]int{0, 1, 2, 4, 6, 7})
	fmt.Println(`result: `, result)
}
