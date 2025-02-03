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

func productExceptSelf(nums []int) []int {
	var result []int = make([]int, len(nums))

	right := make([]int, len(nums))

	products := 1
	for i := len(nums) - 1; i >= 0; i-- {
		products *= nums[i]
		right[i] = products
	}

	products = 1
	for i := 0; i < len(nums)-1; i++ {
		lp := products
		rp := right[i+1]

		result[i] = lp * rp
		products *= nums[i]
	}

	result[len(nums)-1] = products
	return result
}

func main() {
	fmt.Println(`test`)
	result := productExceptSelf([]int{1, 2, 3, 4})

	fmt.Println(result)
}
