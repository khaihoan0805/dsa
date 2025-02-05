package main

import (
	"fmt"
)

// func findMaxAverage(nums []int, k int) float64 {
// 	var max int

// 	for _, v := range nums[:k] {
// 		max += v
// 	}

// 	for i := 1; i < len(nums)-(k-1); i++ {
// 		var temp = 0
// 		slice := nums[i : k+i]
// 		fmt.Println(slice)
// 		for _, v := range slice {
// 			temp += v
// 		}

// 		if temp > max {
// 			max = temp
// 		}
// 	}

// 	var result float64 = float64(max) / float64(k)

// 	return result
// }

func findMaxAverage(nums []int, k int) float64 {
	avg := 0
	for i := 0; i < k; i++ {
		avg += nums[i]
	}

	curr := avg
	for i := k; i < len(nums); i++ {
		curr += nums[i]
		curr -= nums[i-k]
		if curr > avg {
			avg = curr
		}
	}
	return float64(avg) / float64(k)
}

func main() {
	fmt.Println(`test`)
	result := findMaxAverage([]int{0, 1, 1, 3, 3}, 4)

	fmt.Println(result)
}
