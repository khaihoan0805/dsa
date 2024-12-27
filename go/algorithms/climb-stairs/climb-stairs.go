package main

import (
	"fmt"
)

func climbStairsRecursive(methods int, n int) int {
	if n == 0 {
		return methods + 1
	}

	if n >= 1 {
		methods = climbStairsRecursive(methods, n-1)
	}

	if n >= 2 {
		methods = climbStairsRecursive(methods, n-2)
	}

	return methods
}

func climbStairs(n int) int {
	var methods int = 0

	return climbStairsRecursive(methods, n)
}

// func climbStairs(n int) int {
//     next, secondNext := 0, 1
//     for ; n > 0; n-- {
//         next, secondNext = secondNext, next + secondNext
//     }
//     return secondNext
// }

func main() {
	fmt.Println(`test`)

	fmt.Println(climbStairs(5))
}
