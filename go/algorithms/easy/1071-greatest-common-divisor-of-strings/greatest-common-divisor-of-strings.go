package main

import (
	"fmt"
)

// func gcdOfStrings(str1 string, str2 string) string {
// 	var result string = ""

// 	for index := 0; index < len(str1) && index < len(str2); index++ {
// 		temp := str1[0 : index+1]
// 		regex := regexp.MustCompile(fmt.Sprintf("(%s)", temp))

// 		// match1 := regex.MatchString(str1)
// 		// match2 := regex.MatchString(str2)

// 		times1 := regex.FindAllString(str1, -1)
// 		times2 := regex.FindAllString(str2, -1)

// 		// fmt.Println(`time of repeat in 1: `, regex.FindAllString(str1, -1))
// 		// fmt.Println(`time of repeat in 2: `, regex.FindAllString(str2, -1))

// 		if (len(times1) != 0 && len(temp)*len(times1) == len(str1)) && (len(times2) != 0 && len(temp)*len(times2) == len(str2)) {
// 			result = temp
// 		}
// 		// fmt.Println(`result: `, result)
// 	}

// 	return result
// }

func gcdOfStrings(str1 string, str2 string) string {
	if str2+str1 != str1+str2 {
		return ""
	}
	gcdLength := gcd(len(str1), len(str2))
	return str1[:gcdLength]
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(`test`)
	result := gcdOfStrings("ABABAB", "ABAB")
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
