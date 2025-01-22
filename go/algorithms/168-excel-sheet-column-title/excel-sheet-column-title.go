package main

import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	var result string
	for columnNumber > 0 {
		columnNumber--
		divided := columnNumber % 26
		charCode := 'A' + rune(divided)
		result = string(charCode) + result
		columnNumber /= 26
	}
	return result
}

func main() {
	fmt.Println(`test`)
	result := convertToTitle(29)

	fmt.Println(result)
}
