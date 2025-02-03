package main

import (
	"fmt"
	"strconv"
)

// func isPowerOfTwo(n int) bool {
// 	if n == 1 {
// 		return true
// 	}

// 	temp := 1
// 	x := 0

// 	for math.Abs(float64(n)) > math.Abs(float64(temp)) {
// 		temp = int(math.Pow(2, float64(x)))
// 		if temp == n {
// 			return true
// 		}

// 		x++
// 	}

// 	return false

// }

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	fmt.Println(`strconv.FormatInt(int64(n), 2): `, strconv.FormatInt(int64(n), 2))
	fmt.Println(`strconv.FormatInt(int64(n-1), 2): `, strconv.FormatInt(int64(n-1), 2))
	res := n & (n - 1)

	if res == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(`test`)
	result := isPowerOfTwo(8)
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
