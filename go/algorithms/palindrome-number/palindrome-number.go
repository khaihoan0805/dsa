package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPalindrome(x int) bool {
	parsedString := strconv.Itoa(x)

	for i := 0; i <= int(math.Floor(float64(len(parsedString)/2))); i++ {
		// fmt.Println(`------------------------------------`)
		// fmt.Println(`index: `, i)
		// fmt.Println(`parsedString[i]): `, string(parsedString[i]))
		// fmt.Println(`string(parsedString[len(parsedString)-1]: `, string(parsedString[len(parsedString)-i-1]))
		// fmt.Println(`------------------------------------`)
		if string(parsedString[i]) != string(parsedString[len(parsedString)-i-1]) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(100001))
}
