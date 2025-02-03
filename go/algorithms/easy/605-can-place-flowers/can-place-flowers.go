package main

import (
	"fmt"
)

// func canPlaceFlowers(flowerbed []int, n int) bool {
// 	placed := 0

// 	if len(flowerbed) == 1 && flowerbed[0] == 0 {
// 		placed++
// 		return placed >= n
// 	}

// 	if len(flowerbed) == 1 && flowerbed[0] == 1 {
// 		return placed >= n
// 	}

// 	if flowerbed[0] == 0 && flowerbed[1] == 0 {
// 		flowerbed[0] = 1
// 		placed++
// 	}

// 	for index := 1; index < len(flowerbed)-1 && placed < n; index++ {
// 		if flowerbed[index-1] == 0 && flowerbed[index] == 0 && flowerbed[index+1] == 0 {
// 			flowerbed[index] = 1
// 			placed++
// 			index++
// 		}
// 	}

// 	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
// 		flowerbed[len(flowerbed)-1] = 1
// 		placed++
// 	}

// 	return placed >= n
// }

func canPlaceFlowers(flowerbed []int, n int) bool {
	len := len(flowerbed)
	for i := 0; i < len; i++ {
		left := i == 0 || flowerbed[i-1] == 0
		right := i == len-1 || flowerbed[i+1] == 0

		if left && right && flowerbed[i] == 0 {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}

func main() {
	fmt.Println(`test`)
	result := canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1)
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
