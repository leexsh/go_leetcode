package leetcode

/*
题目：用队列实现栈
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
*/

type MyStack struct {
	que1 []int
	que2 []int
}

func ConstructorMyStack() MyStack {
	s := MyStack{
		que1: []int{},
		que2: []int{},
	}
	return s
}

func (this *MyStack) Push(x int) {
	this.que2 = append(this.que2, x)
	for len(this.que1) > 0 {
		this.que2 = append(this.que2, this.que1[0])
		this.que1 = this.que1[1:]
	}
	this.que1, this.que2 = this.que2, this.que1
}

func (this *MyStack) Pop() int {
	v := this.que1[0]
	this.que1 = this.que1[1:]
	return v
}

func (this *MyStack) Top() int {
	return this.que1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.que1) == 0
}
