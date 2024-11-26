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
		// fmt.Println(`----------------------------------------`)
		// fmt.Println(`open: `, open)
		if slices.Contains(openChars, string(s[index])) {
			// fmt.Println(`open Char: `, string(s[index]))
			open = append(open, string(s[index]))
			// fmt.Println(`add to open: `, open)
			continue
		}

		if slices.Contains(closeChars, string(s[index])) {
			// fmt.Println(`close Char: `, string(s[index]))
			if len(openChars) < 1 {
				return false
			}
			tail := string(open[len(open)-1])
			// fmt.Println(`tail: `, tail)

			validClose := openClose[tail]
			// fmt.Println(`validClose: `, validClose == string(s[index]))

			if validClose == string(s[index]) {
				open = open[:len(open)-1]
				// fmt.Println(`remove to open: `, open)
			} else {
				return false
			}
		}
		// fmt.Println(`----------------------------------------`)
	}

	// fmt.Println(`open: `, open)

	// fmt.Println(`close: `, close)

	// if len(open) != len(close) {
	// 	return false
	// }

	// for i := 0; i < len(open); i++ {
	// 	value := openClose[string(open[len(open)-1])]
	// 	fmt.Println(`string(open[len(open)-1]): `, string(open[len(open)-1]))
	// 	fmt.Println(`value: `, value)
	// 	fmt.Println(`close[i]: `, close[i])
	// 	if value != close[i] {
	// 		fmt.Println("Banana exists with value:", value)
	// 		return false
	// 	}
	// }

	if len(open) > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(`test`)
	fmt.Println(isValid("([)]{}"))
}
