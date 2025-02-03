package main

import (
	"fmt"
)

type MyStack struct {
	deque []int
}

func (this *MyStack) Push(x int) {
	this.deque = append([]int{x}, this.deque...)
}

func (this *MyStack) Pop() int {
	value := this.deque[0]

	this.deque = this.deque[1:]

	return value
}

func (this *MyStack) Top() int {
	return this.deque[0]
}

func (this *MyStack) Empty() bool {
	return len(this.deque) == 0
}

type MyQueue struct {
	stack1 MyStack
	stack2 MyStack
}

func Constructor() MyQueue {
	return MyQueue{stack1: MyStack{}, stack2: MyStack{}}
}

func (this *MyQueue) Push(x int) {

	this.stack1.Push(x)

}

func (this *MyQueue) Pop() int {
	for this.stack1.Empty() == false {
		this.stack2.Push(this.stack1.Pop())
	}

	value := this.stack2.Pop()

	for this.stack2.Empty() == false {
		this.stack1.Push(this.stack2.Pop())
	}

	return value
}

func (this *MyQueue) Peek() int {
	fmt.Println(`this.stack1.deque before:`, this.stack1.deque)

	for this.stack1.Empty() == false {
		this.stack2.Push(this.stack1.Pop())
	}

	fmt.Println(`this.stack2.deque:`, this.stack2.deque)

	value := this.stack2.Top()

	for this.stack2.Empty() == false {
		this.stack1.Push(this.stack2.Pop())
	}

	fmt.Println(`this.stack1.deque:`, this.stack1.deque)

	return value
}

func (this *MyQueue) Empty() bool {
	return this.stack1.Empty()
}

func main() {
	fmt.Println(`test`)
	queue := Constructor()

	queue.Push(1)
	queue.Push(2)

	fmt.Println(`queue.Peek():`, queue.Peek())
	fmt.Println(`queue.pop():`, queue.Pop())
	// fmt.Println(math.Pow(-2, 3))

	// fmt.Println(result)
}
