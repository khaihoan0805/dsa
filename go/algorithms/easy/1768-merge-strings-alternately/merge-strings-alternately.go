package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	var result string = ""
	// index := 0
	// for index < len(word1) && index < len(word2) {
	// 	result += string(word1[index]) + string(word2[index])
	// 	index++
	// }

	// if len(word1) > len(word2) {
	// 	remaining := word1[index:]
	// 	result += remaining
	// }

	// if len(word2) > len(word1) {
	// 	remaining := word2[index:]
	// 	result += remaining
	// }

	// for index := 0; index < len(word1) || index < len(word2); index++ {
	// 	charA := ""
	// 	if index <= len(word1)-1 {
	// 		charA = string(word1[index])
	// 	}

	// 	charB := ""
	// 	if index <= len(word2)-1 {
	// 		charB = string(word2[index])
	// 	}

	// 	result += charA + charB
	// }

	for len(word1) != 0 && len(word2) != 0 {
		result += string(word1[0]) + string(word2[0])
		word1 = word1[1:]
		word2 = word2[1:]
	}

	if len(word1) > 0 {
		result += word1
	}

	if len(word2) > 0 {
		result += word2
	}

	return result
}

func main() {
	fmt.Println(`test`)
	result := mergeAlternately("ab", "pqrs")
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
