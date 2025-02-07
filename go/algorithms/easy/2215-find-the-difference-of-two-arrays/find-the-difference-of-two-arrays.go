package main

import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for _, v := range nums1 {
		map1[v] = 1
	}

	for _, v := range nums2 {
		map2[v] = 1
	}

	var distinct1 []int
	for key := range map1 {
		if _, exists := map2[key]; !exists {
			distinct1 = append(distinct1, key)
		}
	}

	var distinct2 []int
	for key := range map2 {
		if _, exists := map1[key]; !exists {
			distinct2 = append(distinct2, key)
		}
	}

	return [][]int{distinct1, distinct2}
}

func main() {
	fmt.Println(`test`)
	result := findDifference([]int{-5, 1, 5, 0, -7}, []int{-5, 1, 5, 0, -7})

	fmt.Println(result)
}
