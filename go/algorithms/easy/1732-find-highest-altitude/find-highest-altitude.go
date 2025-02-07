package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	max, pre := 0, 0

	for _, v := range gain {
		pre += v
		if pre > max {
			max = pre
		}
	}

	return max
}

func main() {
	fmt.Println(`test`)
	result := largestAltitude([]int{-5, 1, 5, 0, -7})

	fmt.Println(result)
}
