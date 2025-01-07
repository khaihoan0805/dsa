package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var res int = 0

	for index := 0; index < len(nums); index++ {
		res ^= nums[index]
	}

	return res
}

func main() {
	fmt.Println(`test`)

	fmt.Println(singleNumber([]int{2, 2, 1}))
}
