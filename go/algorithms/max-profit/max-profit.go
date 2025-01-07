package main

import (
	"fmt"
	"math"
)

// func maxProfit(prices []int) int {
// 	// result := make([]int, len(prices))
// 	max := 0

// 	for index := 0; index <= len(prices)-1; index++ {
// 		days := prices[index+1:]

// 		for i := 0; i < len(days); i++ {
// 			profit := days[i] - prices[index]

// 			if profit > max {
// 				max = profit
// 			}
// 		}
// 	}

// 	return max
// }

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, currentPrice := range prices {
		minPrice = min(currentPrice, minPrice)
		currentProfit := currentPrice - minPrice
		maxProfit = max(maxProfit, currentProfit)
	}

	return maxProfit
}

func main() {
	fmt.Println(`test`)

	// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{2, 7, 5, 4, 1, 6, 3}))
}
