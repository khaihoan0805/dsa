package main

import (
	"fmt"
)

func mySqrt(x int) int {
	var result int = 0

	for {
		if (result * result) > x {
			return result - 1
		}

		result++
	}
}

func main() {
	fmt.Println(`test`)

	fmt.Println(mySqrt(4))
}
