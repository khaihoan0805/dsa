package main

import (
	"fmt"
	"strings"
)

// func isVowel(c string) bool {
// 	vowels := "aeiouAEIOU"

// 	for _, v := range vowels {
// 		if strings.ToLower(c) == strings.ToLower(v) {
// 			return true
// 		}
// 	}

// 	return false
// }

func reverseWords(s string) string {
	temp := ""

	result := ""

	for index := len(s) - 1; index >= 0; index-- {
		if string(s[index]) == " " {
			if len(result) > 0 && len(temp) > 0 {
				result = result + " " + temp
			} else {
				result += temp
			}
			temp = ""
		} else {
			temp = string(s[index]) + temp
		}
	}

	if len(result) > 0 && len(temp) > 0 {
		result = result + " " + temp
	} else {
		result += temp
	}

	return strings.TrimSpace(result)
}

func main() {
	fmt.Println(`test`)
	result := reverseWords("a good   example")
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
