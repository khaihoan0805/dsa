package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)

		for j := range result[i] {
			result[i][j] = 1
		}
	}

	for i := range result {
		if i == 0 || i == 1 {
			continue
		}

		for j := 0; j < len(result[i]); j++ {
			if j == 0 || j == len(result[i])-1 {
				continue
			}

			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}

func main() {
	fmt.Println(`test`)

	fmt.Println(generate(5))
}
