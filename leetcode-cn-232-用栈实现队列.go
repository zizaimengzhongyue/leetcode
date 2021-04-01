package main

import (
	"fmt"
)

type Stack struct {
	s []int
}

func (this *Stack) Push(x int) {
	this.s = append(this.s, x)
}

func (this *Stack) Pop() int {
	res := this.s[len(this.s)-1]
	this.s = this.s[0 : len(this.s)-1]
	return res
}

func (this *Stack) Size() int {
	return len(this.s)
}

func (this *Stack) Peek() int {
	return this.s[len(this.s)-1]
}

func NewStack() Stack {
	return Stack{}
}

type MyQueue struct {
	s1 Stack
	s2 Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{s1: NewStack(), s2: NewStack()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.swap(0)
	res := this.s2.Pop()
	this.swap(1)
	return res
}

func (this *MyQueue) swap(flag int) {
	if flag == 0 {
		ln := this.s1.Size()
		for i := 0; i < ln; i++ {
			this.s2.Push(this.s1.Pop())
		}
	} else {
		ln := this.s2.Size()
		for i := 0; i < ln; i++ {
			this.s1.Push(this.s2.Pop())
		}
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.swap(0)
	res := this.s2.Peek()
	this.swap(1)
	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.Size() == 0
}

func main() {
	q := MyQueue{}
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}
