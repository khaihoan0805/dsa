package main

import (
	"fmt"
)

type MyStack struct {
	queue []string
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x string) {
	this.queue = append([]string{x}, this.queue...)
}

func (this *MyStack) Pop() string {
	if this.Empty() {
		return ""
	}

	value := this.queue[0]

	this.queue = this.queue[1:]

	return value
}

func (this *MyStack) Top() string {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func removeStars(s string) string {
	var stack MyStack = MyStack{}

	for _, v := range s {
		if string(v) != "*" {
			stack.Push(string(v))
		} else {
			stack.Pop()
		}
	}

	result := ""
	for !stack.Empty() {
		result = stack.Pop() + result
	}

	return result
}

func main() {
	fmt.Println(`test`)
	result := removeStars("leet**cod*e")
	// fmt.Println(math.Pow(-2, 3))

	fmt.Println(result)
}
