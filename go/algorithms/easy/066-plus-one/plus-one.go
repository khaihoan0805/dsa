package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	var result = make([]int, len(digits)+1, len(digits)+1)
	fmt.Println(`len of result: `, len(result))

	var isNeedToPlusOne int = 1
	for index := len(digits) - 1; index >= 0; index-- {
		plusOne := digits[index] + isNeedToPlusOne

		if plusOne >= 10 {
			isNeedToPlusOne = 1
		} else {
			isNeedToPlusOne = 0
		}

		result[index+1] = plusOne % 10
	}

	if isNeedToPlusOne == 1 {
		result[0] = 1
		return result
	}

	return result[1:]
}

func main() {
	fmt.Println(`test`)
	fmt.Println(`final: `, plusOne([]int{1, 2, 3}))
}
