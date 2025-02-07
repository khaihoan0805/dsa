package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	var left, right []int = make([]int, len(nums)), make([]int, len(nums))

	leftTotal := 0
	for i := 1; i < len(nums); i++ {
		leftTotal += nums[i-1]
		left[i] = leftTotal
	}

	rightTotal := 0
	for i := len(nums) - 2; i >= 0; i-- {
		rightTotal += nums[i+1]
		right[i] = rightTotal
	}

	for i := 0; i < len(nums); i++ {
		if left[i] == right[i] {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(`test`)
	result := pivotIndex([]int{-1, -1, 0, 1, 1, 0})

	fmt.Println(result)
}
