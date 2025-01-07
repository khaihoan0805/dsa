package main

import (
	"fmt"
	"unicode"
)

// func isPalindrome(s string) bool {
// 	var nonAlphanumericString string = ""

// 	for _, char := range s {
// 		if unicode.IsLetter(char) || unicode.IsDigit(char) {
// 			nonAlphanumericString += string(unicode.ToLower(char))
// 		}
// 	}

// 	if nonAlphanumericString == "" {
// 		return true
// 	}

// 	fmt.Println("string: ", nonAlphanumericString)

// 	for index := 0; index < len(nonAlphanumericString)/2; index++ {
// 		if nonAlphanumericString[index] != nonAlphanumericString[len(nonAlphanumericString)-1-index] {
// 			return false
// 		}
// 	}

// 	return true
// }

func isPalindrome(s string) bool {
	var head, tail int = 0, len(s) - 1

	for head <= tail {
		if !unicode.IsLetter(rune(s[head])) && !unicode.IsDigit(rune(s[head])) {
			head += 1
			continue
		}

		if !unicode.IsLetter(rune(s[tail])) && !unicode.IsDigit(rune(s[tail])) {
			tail -= 1
			continue
		}

		if unicode.ToLower(rune(s[head])) != unicode.ToLower(rune(s[tail])) {
			return false
		}

		head++
		tail--
	}

	return true
}

func main() {
	fmt.Println(`test`)

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
