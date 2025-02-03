package main

import (
	"fmt"
)

// func majorityElement(nums []int) []int {
// 	temp := make(map[int]int)

// 	result := []int{}

// 	for _, v := range nums {
// 		value, exists := temp[v]

// 		if !exists {
// 			temp[v] = 1
// 		} else {
// 			if temp[v] == -1 {
// 				continue
// 			}

// 			temp[v] = value + 1

// 			if temp[v] > len(nums)/3 {
// 				result = append(result, v)
// 				temp[v] = -1
// 			}

// 		}

// 		// if temp[v] > len(nums)/3 {
// 		// 	result = append(result, v)
// 		// }
// 	}

// 	return result
// }

func majorityElement(nums []int) []int {
	temp := make(map[int]int)

	result := []int{}

	for _, v := range nums {
		value, exists := temp[v]

		if !exists {
			temp[v] = 1
		} else {
			if temp[v] == -1 {
				continue
			}

			temp[v] = value + 1
		}

		if temp[v] != -1 && temp[v] > len(nums)/3 {
			result = append(result, v)
			temp[v] = -1
		}
	}

	return result
}

func main() {
	fmt.Println(`test`)
	result := majorityElement([]int{2, 2, 2, 1, 1, 1, 5})

	fmt.Println(result)
}
