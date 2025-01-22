package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	patternS := make(map[string]int)
	patternT := make(map[string]int)

	for index := 0; index < len(s); index++ {
		// vS, okS := patternS[string(s[index])]
		// // vT, okT := patternT[string(t[index])]
		// fmt.Printf("patternS[%s]: %d \n", string(s[index]), patternT[(string(s[index]))])
		// fmt.Printf("patternT[%s]: %d \n", string(t[index]), patternT[(string(t[index]))])

		if patternS[string(s[index])] != patternT[(string(t[index]))] {
			return false
		} else {
			patternS[string(s[index])] = index + 1
			patternT[string(t[index])] = index + 1
		}
	}

	return true
}

// func isIsomorphic(s string, t string) bool {
// 	sPattern, tPattern := map[uint8]int{}, map[uint8]int{}
// 	for index := range s {
// 		fmt.Println(`sPattern[s[index]]: `, sPattern[s[index]])
// 		fmt.Println(`tPattern[t[index]]: `, tPattern[t[index]])
// 		if sPattern[s[index]] != tPattern[t[index]] {
// 			return false
// 		} else {
// 			sPattern[s[index]] = index + 1
// 			tPattern[t[index]] = index + 1
// 		}
// 	}

// 	return true
// }

// func isIsomorphic(s string, t string) bool {
// 	map1 := make([]int, 128) // Stores frequency of s
// 	map2 := make([]int, 128) // Stores frequency of t

// 	for i := 0; i < len(s); i++ {
// 		sCh := s[i]
// 		tCh := t[i]

// 		if map1[sCh] == 0 && map2[tCh] == 0 {
// 			map1[sCh] = int(tCh)
// 			map2[tCh] = int(sCh)
// 		} else if map1[sCh] != int(tCh) || map2[tCh] != int(sCh) {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	fmt.Println(`test`)
	// Define a map
	result := isIsomorphic("aa", "ab")
	fmt.Println(`result: `, result)
}
