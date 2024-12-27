package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	a := []int{}
	indices := []int{}

	for i, v := range nums {
		isExisted := false
		for j := 0; j < len(a); j++ {
			if v == a[j] {
				isExisted = true
				break
			}
		}

		if isExisted == true {
			indices = append(indices, i)
		} else {
			a = append(a, v)
		}
	}

	// fmt.Println(`a: `, a)
	// fmt.Println(`indices: `, indices)

	minus := 0
	for _, value := range indices {
		nums = append(nums[:value-minus], nums[value+1-minus:]...)
		fmt.Println(`nums: `, nums)
		minus++
	}

	return len(a)
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Using two pointers approach
	// i points to the last unique element
	// j iterates through the array
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func main() {
	fmt.Println(`test`)
	nums := []int{0, 0, 1, 1, 1, 2, 2}

	removeDuplicates1(nums)
	fmt.Println(`nums: `, nums)
}
