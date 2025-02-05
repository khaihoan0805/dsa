package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxArea := 0
	i, j := 0, len(height)-1

	for i < j {
		curWidth := j - i
		curHeight := min(height[i], height[j])
		maxArea = max(maxArea, curWidth*curHeight)

		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(`test`)
	test := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(test)
	// result := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	// result := increasingTriplet([]int{1, 2, 3, 4, 5})
	// result := increasingTriplet([]int{1, 2, 2147483647})

	fmt.Println(result)
}
