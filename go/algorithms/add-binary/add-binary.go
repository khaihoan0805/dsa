package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var result string = ""
	var remember int = 0

	var longer string = a
	var shorter string = b

	if len(longer) < len(shorter) {
		longer = b
		shorter = a
	}

	var gap int = len(longer) - len(shorter)

	for index := len(longer) - 1; index >= 0; index-- {
		var temp string = result
		aIndex, _ := strconv.ParseInt(string(longer[index]), 10, 64)
		var bIndex int64 = 0

		if index-gap >= 0 {
			bIndex, _ = strconv.ParseInt(string(shorter[index-gap]), 10, 64)
		}

		var plus int = int(aIndex) + int(bIndex) + remember

		if plus < 2 {
			remember = 0
		} else {
			remember = 1
		}

		result = strconv.Itoa(plus%2) + temp
	}

	if remember == 1 {
		return "1" + result
	}

	return result
}

func main() {
	fmt.Println(`test`)
	var b string = "10011"
	var a string = "1101"

	fmt.Println(addBinary(a, b))
	value1, _ := strconv.ParseInt(a, 2, 64)
	// fmt.Println(`value1: `, value1)

	value2, _ := strconv.ParseInt(b, 2, 64)

	fmt.Println(strconv.FormatInt(value1+value2, 2))
}
