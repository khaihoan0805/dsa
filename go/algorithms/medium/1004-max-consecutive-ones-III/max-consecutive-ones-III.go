package main

import (
	"fmt"
)

func findLastIndex(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1 // Return -1 if the target is not found
}

func longestOnes(nums []int, k int) int {
	if k == 0 {
		result := 0
		total := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == 1 {
				total++
				if total > result {
					result = total
				}
			} else {
				total = 0
			}
		}

		return result
	}

	last0Index := findLastIndex(nums, 0)

	if last0Index == -1 {
		return len(nums)
	}

	index := last0Index + 1

	count := 1
	slice := nums[:index]
	max := len(slice)
	for index < len(nums) {
		if nums[index] == 1 {
			slice = append(slice, 1)
		}

		if nums[index] == 0 {
			count++
			if count <= k {
				slice = append(slice, 0)
			} else {
				slice = nums[last0Index+1 : index+1]
				last0Index += findLastIndex(slice, 0) + 1
				count--
			}
		}

		if len(slice) > max {
			max = len(slice)
		}

		index++
	}

	return max
}

func main() {
	fmt.Println(`test`)
	// result := longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	// result := longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	result := longestOnes([]int{0, 0, 1, 1, 1, 0, 0}, 0)

	fmt.Println(result)
	// fmt.Println(result1)
}
