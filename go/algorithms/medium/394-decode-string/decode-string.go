package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	openClose := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	var closeChars = []string{"}", "]", ")"}

	var stack []string
	subString := ""
	for _, v := range s {
		char := string(v)
		if !slices.Contains(closeChars, char) {
			stack = append(stack, char)
		} else {
			for stack[len(stack)-1] != openClose[char] {
				subString = stack[len(stack)-1] + subString
				stack = stack[:len(stack)-1]
			}

			stack = stack[:len(stack)-1]
			numString := ""
			for len(stack) > 0 {
				_, err := strconv.Atoi(string(stack[len(stack)-1] + numString))
				if err == nil {
					numString = stack[len(stack)-1] + numString
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}

			multiplier, _ := strconv.Atoi(numString)
			subString = strings.Repeat(subString, multiplier)
		}

		for _, c := range subString {
			stack = append(stack, string(c))
		}
		subString = ""
	}

	return strings.Join(stack, "")
}

func main() {
	fmt.Println(`test`)
	result := decodeString("100[leetcode]")
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
