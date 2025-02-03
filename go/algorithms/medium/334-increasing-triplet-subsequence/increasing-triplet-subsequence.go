package main

import (
	"fmt"
)

// func productExceptSelf(nums []int) []int {
// 	var result []int = make([]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		subResult := 1
// 		for j := 0; j < len(nums); j++ {
// 			if j == i {
// 				continue
// 			}

// 			subResult *= nums[j]
// 		}

// 		result[i] = subResult
// 	}

// 	return result
// }

// func increasingTriplet(nums []int) bool {
// 	remind := nums[0]

// 	for i := 0; i < len(nums)-3; i++ {
// 		if nums[i] > remind {
// 			continue
// 		}

// 		subi := nums[1:]
// 		for j := 0; j < len(subi) - 2; j++ {

// 		}
// 	}
// }

func increasingTriplet(nums []int) bool {
	var ls []bool = make([]bool, len(nums))
	ls[0] = false
	smallest := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > smallest {
			ls[i] = true
		} else {
			ls[i] = false
			smallest = nums[i]
		}
	}

	var rg []bool = make([]bool, len(nums))
	rg[len(nums)-1] = false
	greatest := nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		if greatest > nums[i] {
			rg[i] = true
		} else {
			rg[i] = false
			greatest = nums[i]
		}
	}

	for i := 1; i < len(nums)-1; i++ {
		if ls[i] == true && rg[i] == true {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(`test`)
	result := increasingTriplet([]int{6, 7, 1, 2})
	// result := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	// result := increasingTriplet([]int{1, 2, 3, 4, 5})
	// result := increasingTriplet([]int{1, 2, 2147483647})

	fmt.Println(result)
}
