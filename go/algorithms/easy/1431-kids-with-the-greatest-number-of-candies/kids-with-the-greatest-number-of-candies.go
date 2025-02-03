package main

import (
	"fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatestBeforeExtra := 0

	for _, v := range candies {
		if v > greatestBeforeExtra {
			greatestBeforeExtra = v
		}
	}

	var result []bool = make([]bool, len(candies))
	for index := 0; index < len(candies); index++ {
		result[index] = candies[index]+extraCandies >= greatestBeforeExtra
	}

	return result
}

func main() {
	fmt.Println(`test`)
	result := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
