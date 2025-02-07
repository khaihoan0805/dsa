package main

import (
	"fmt"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var map1 map[rune]int = make(map[rune]int)
	for _, v := range word1 {
		if value, exists := map1[v]; exists {
			map1[v] = value + 1
		} else {
			map1[v] = 1
		}
	}

	var map2 map[rune]int = make(map[rune]int)
	for _, v := range word2 {
		if value, exists := map2[v]; exists {
			map2[v] = value + 1
		} else {
			map2[v] = 1
		}
	}

	var concurrence1 map[int]int = make(map[int]int)
	for _, value := range map1 {
		if current, exists := concurrence1[value]; exists {
			concurrence1[value] = current + 1
		} else {
			concurrence1[value] = 1
		}
	}

	var concurrence2 map[int]int = make(map[int]int)
	for key, value := range map2 {
		if _, exists := map1[key]; !exists {
			return false
		}

		if current, exists := concurrence2[value]; exists {
			concurrence2[value] = current + 1
		} else {
			concurrence2[value] = 1
		}
	}

	for key := range concurrence1 {
		if concurrence1[key] != concurrence2[key] {
			return false
		}
	}

	return true
}

type Solution struct{}

func (s *Solution) CloseStrings(word1 string, word2 string) bool {
	// Check if the lengths are equal
	length := len(word1)
	if length != len(word2) {
		return false
	}

	// Initialize arrays for character frequencies and cumulative frequencies
	chars1 := make([]int, 26)
	chars2 := make([]int, 26)
	values1 := make([]int, length+1)
	values2 := make([]int, length+1)

	// Process word1 and word2 to update the frequency arrays
	s.process(chars1, values1, word1)
	s.process(chars2, values2, word2)

	// Check if characters present in word1 are also present in word2
	for i := 0; i < 26; i++ {
		if chars1[i] > 0 && chars2[i] == 0 {
			return false
		}
	}

	// Check if cumulative frequencies are equal for both words
	for i := 0; i < length; i++ {
		if values1[i] != values2[i] {
			return false
		}
	}

	// If all checks pass, the words can be made equivalent
	return true
}

// Helper method to update character and cumulative frequencies
func (s *Solution) process(chars []int, values []int, word string) {
	for _, c := range word {
		chars[c-'a']++
		values[chars[c-'a']]++
	}
}

func main() {
	fmt.Println(`test`)
	result := closeStrings("abc", "bca")

	var solution Solution = Solution{}
	solution.CloseStrings("abc", "bca")
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
