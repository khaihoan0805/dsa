package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var prefix string = ""

	if len(strs) < 1 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	if len(strs[0]) == 0 {
		return ""
	}

	var rootEl string = strs[0]

	for i := range rootEl {
		var forCompared string = rootEl[:i+1]

		for _, value := range strs {
			if i+1 > len(value) {
				return prefix
			}

			if forCompared != value[0:i+1] {
				return prefix
			}
		}

		prefix = forCompared
	}
	return prefix
}

func main() {
	fmt.Println(`test`)
	test := longestCommonPrefix([]string{"flower", "flow", "flight"})

	fmt.Println(test)
}
