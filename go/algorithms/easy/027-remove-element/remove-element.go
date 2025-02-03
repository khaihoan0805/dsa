package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	isNotEqualElements := 0
	indices := []int{}
	for i, v := range nums {
		if v == val {
			indices = append(indices, i)
		} else {
			isNotEqualElements++
		}
	}

	// fmt.Println(`indices: `, indices)
	// fmt.Println(`isNotEqualElements: `, isNotEqualElements)

	minus := 0
	for _, value := range indices {
		nums = append(nums[:value-minus], nums[value+1-minus:]...)
		// fmt.Println(`nums: `, nums)
		minus++
	}

	return isNotEqualElements
}

func main() {
	fmt.Println(`test`)
	nums := []int{0, 0, 1, 1, 1, 2, 2}

	removeElement(nums, 1)
	fmt.Println(`nums: `, nums)
}
