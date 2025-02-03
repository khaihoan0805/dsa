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

func reverseVowels(s string) string {
	if len(s) < 2 {
		return s
	}

	runes := []rune(s)
	vowels := "aeiouAEIOU"

	left := 0
	right := len(s) - 1
	for left < right {
		isVowelLeft := strings.ContainsRune(vowels, runes[left])
		isVowelRight := strings.ContainsRune(vowels, runes[right])

		if isVowelLeft && isVowelRight {
			temp := runes[left]
			runes[left] = runes[right]
			runes[right] = temp

			left++
			right--
			continue
		}

		if !isVowelRight {
			right--
			continue
		}

		if !isVowelLeft {
			left++
			continue
		}
	}

	return string(runes)
}

func main() {
	fmt.Println(`test`)
	result := reverseVowels("IceCreAm")
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
