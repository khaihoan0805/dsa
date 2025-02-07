package main

import (
	"fmt"
)

func uniqueOccurrences(arr []int) bool {
	var occurences map[int]int = make(map[int]int)
	var count map[int]bool = make(map[int]bool)

	for _, v := range arr {
		if value, exists := occurences[v]; exists {
			occurences[v] = value + 1
		} else {
			occurences[v] = 1
		}
	}

	for _, value := range occurences {
		if _, exists := count[value]; exists {
			return false
		} else {
			count[value] = true
		}
	}

	return true
}

func main() {
	fmt.Println(`test`)
	result := uniqueOccurrences([]int{1, 2})

	fmt.Println(result)
}
