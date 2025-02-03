package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	reminder := make(map[int]int)

	for _, v := range nums {
		if reminder[v] == 1 {
			return true
		}

		reminder[v] += 1
	}

	return false
}

func main() {
	fmt.Println(`test`)

	result := containsDuplicate([]int{1, 2, 3, 1})

	fmt.Println(`result: `, result)
}
