package main

import (
	"fmt"
	"math"
)

func titleToNumber(columnTitle string) int {
	var result int = 0
	var multiplier float64 = 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		charCodeNumber := rune(columnTitle[i] - 'A')
		result += (int(charCodeNumber) + 1) * int(math.Pow(26, multiplier))
		multiplier++
	}

	return result
}

func main() {
	fmt.Println(`test`)
	result := titleToNumber("AB")

	fmt.Println(result)
}
