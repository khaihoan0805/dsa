package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	for index := 0; index <= len(haystack)-len(needle); index++ {
		if needle[0] == haystack[index] && len(haystack)-index >= len(needle) {
			if haystack[index:index+len(needle)] == needle {
				return index
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(`test`)
	fmt.Println(strStr("sadbutsad", "sad"))
}
