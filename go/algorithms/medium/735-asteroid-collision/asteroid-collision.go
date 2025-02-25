package main

import (
	"fmt"
	"math"
)

// type MyStack struct {
// 	queue []int
// }

// func Constructor() MyStack {
// 	return MyStack{}
// }

// func (this *MyStack) Push(x int) {
// 	this.queue = append([]int{x}, this.queue...)
// }

// func (this *MyStack) Pop() int {
// 	value := this.queue[0]

// 	this.queue = this.queue[1:]

// 	return value
// }

// func (this *MyStack) Top() int {
// 	return this.queue[0]
// }

// func (this *MyStack) Empty() bool {
// 	return len(this.queue) == 0
// }

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) == 0 {
		return asteroids
	}

	var stack []int

	for _, v := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}

		isDeleted := false
		for (len(stack) != 0 && ((stack[len(stack)-1] > 0) && (v < 0))) && !isDeleted {
			if math.Abs(float64(v)) > math.Abs(float64(stack[len(stack)-1])) {
				stack = stack[:len(stack)-1]
			} else if math.Abs(float64(v)) == math.Abs(float64(stack[len(stack)-1])) {
				stack = stack[:len(stack)-1]
				isDeleted = true
			} else {
				isDeleted = true
			}
		}

		if !isDeleted {
			stack = append(stack, v)
		}
	}

	return stack
}

func main() {
	fmt.Println(`test`)
	result := asteroidCollision([]int{-2, -1, 1, 2})

	fmt.Println(result)
}
