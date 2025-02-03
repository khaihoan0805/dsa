package main

import (
	"fmt"
	"slices"
)

func isValid(s string) bool {
	openClose := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	var openChars = []string{"{", "[", "("}
	var closeChars = []string{"}", "]", ")"}
	open := []string{}

	for index := 0; index < len(s); index++ {
		if slices.Contains(openChars, string(s[index])) {
			open = append(open, string(s[index]))
			continue
		}

		if slices.Contains(closeChars, string(s[index])) {
			if len(openChars) < 1 {
				return false
			}
			tail := string(open[len(open)-1])

			validClose := openClose[tail]

			if validClose == string(s[index]) {
				open = open[:len(open)-1]
			} else {
				return false
			}
		}
	}

	if len(open) > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(`test`)
	fmt.Println(isValid("([)]{}"))
}
