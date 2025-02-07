package main

import (
	"fmt"
)

// func equalPairs(grid [][]int) int {
// 	rows := make(map[string]int)
// 	columns := make(map[string]int)

// 	for i := 0; i < len(grid); i++ {
// 		row := ""
// 		column := ""
// 		for j := 0; j < len(grid[i]); j++ {
// 			row += strconv.Itoa(grid[i][j]) + " "
// 			column += strconv.Itoa(grid[j][i]) + " "
// 		}

// 		if v, exists := rows[row]; exists {
// 			rows[row] = v + 1
// 		} else {
// 			rows[row] = 1
// 		}

// 		if v, exists := columns[column]; exists {
// 			columns[column] = v + 1
// 		} else {
// 			columns[column] = 1
// 		}
// 	}

// 	total := 0

// 	for key, value := range columns {
// 		if valueRow, exists := rows[key]; exists {
// 			total += (valueRow * value)
// 		}
// 	}

// 	return total
// }

func equalPairs(grid [][]int) int {
	n := len(grid)
	m := make(map[[200]int]int)
	arr := [200]int{}
	for i := 0; i < n; i++ {
		copy(arr[:], grid[i])
		m[arr]++
	}
	res := 0

	for i := 0; i < n; i++ {
		arr = [200]int{}
		for j := 0; j < n; j++ { // copy column to arr
			arr[j] = grid[j][i]
		}
		if v, ok := m[arr]; ok {
			res += v
		}
	}
	return res
}

func main() {
	fmt.Println(`test`)
	// test := []int{3, 1, 3, 4, 5, 3}
	result := equalPairs(
		[][]int{
			{3, 1, 2, 2},
			{1, 4, 4, 5},
			{2, 4, 2, 2},
			{2, 4, 2, 2},
		},
	)
	// result := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	// result := increasingTriplet([]int{1, 2, 3, 4, 5})
	// result := increasingTriplet([]int{1, 2, 2147483647})

	fmt.Println(result)
}
