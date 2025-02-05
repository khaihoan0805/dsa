package main

import (
	"fmt"
)

// func moveZeroes(nums []int) {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			for j := i + 1; j < len(nums); j++ {
// 				if nums[j] != 0 {
// 					nums[i] = nums[j]
// 					nums[j] = 0
// 					break
// 				}
// 			}
// 		}
// 	}

// }

func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}

func main() {
	fmt.Println(`test`)
	test := []int{0, 1, 0, 3, 12}
	moveZeroes(test)
	// result := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	// result := increasingTriplet([]int{1, 2, 3, 4, 5})
	// result := increasingTriplet([]int{1, 2, 2147483647})

	fmt.Println(test)
}
