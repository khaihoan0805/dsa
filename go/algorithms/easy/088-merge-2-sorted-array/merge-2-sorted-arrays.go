package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var result []int = append(nums1[:m], nums2[:n]...)
	sort.Ints(result)

	return result
}

// func merge(nums1 []int, m int, nums2 []int, n int) []int {
// 	var result []int = []int{}

// 	var pivotM, pivotN int = 0, 0

// 	for pivotM <= m-1 && pivotN <= n-1 {

// 		var valueM = nums1[pivotM]

// 		var valueN = nums2[pivotN]

// 		if valueM <= valueN {
// 			result = append(result, valueM)
// 			pivotM++
// 		} else {
// 			result = append(result, valueN)
// 			pivotN++
// 		}

// 	}

// 	if pivotM <= m-1 && len(nums1) > 0 {
// 		var remember int = 1

// 		if len(nums1) == 1 {
// 			remember = 0
// 		}

// 		result = append(result, nums1[pivotM:(m-pivotM)+remember]...)
// 	}

// 	if pivotN <= n-1 && len(nums2) > 0 {
// 		var remember int = 1

// 		if len(nums2) == 1 {
// 			remember = 0
// 		}
// 		result = append(result, nums2[pivotN:(n-pivotN)+remember]...)
// 	}

// 	return result
// }

func main() {
	fmt.Println(`test`)

	// fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))

	// fmt.Println(merge([]int{1}, 1, []int{}, 0))

	fmt.Println(merge([]int{0}, 0, []int{1}, 1))
}
