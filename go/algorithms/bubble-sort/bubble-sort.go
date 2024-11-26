package main

import (
	"fmt"
)

func bubbleSort(array []int) []int {
	var arr []int = array
	var temp int
	for outer := len(arr); outer >= 1; outer-- {
		fmt.Println(`outer: `, outer)
		for inner := 0; inner < outer; inner++ {
			if int(arr[inner]) > int(arr[inner+1]) {
				fmt.Println(array)
				temp = array[inner]
				arr[inner] = array[inner+1]
				arr[inner] = temp
			}

		}
	}

	return array
}

func main() {
	fmt.Println()
}
