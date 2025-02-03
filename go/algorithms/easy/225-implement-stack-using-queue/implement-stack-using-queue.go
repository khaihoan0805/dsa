package main

import (
	"fmt"
)

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue = append([]int{x}, this.queue...)
}

func (this *MyStack) Pop() int {
	value := this.queue[0]

	this.queue = this.queue[1:]

	return value
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	fmt.Println(`test`)
	// root := createTreeFromArray([]int{1, 2, 3, 4, 5, 6})
	// var result int = countNodes(root)
	// fmt.Println(result)
}
