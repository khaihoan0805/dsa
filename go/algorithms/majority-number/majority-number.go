package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	temp := make(map[int]int)

	result := 0

	for _, v := range nums {
		value, exists := temp[v]

		if !exists {
			temp[v] = 1
		} else {
			temp[v] = value + 1
		}

		if temp[v] > len(nums)/2 {
			return v
		}
	}

	return result
}

func main() {
	fmt.Println(`test`)
	result := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})

	fmt.Println(result)
}
