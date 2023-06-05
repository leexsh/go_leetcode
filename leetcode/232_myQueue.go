package leetcode

/*
題目：用栈实现队列
*/
type MyQueue struct {
	que1 []int
	que2 []int
}

func ConstructorMyque() MyQueue {
	return MyQueue{
		que1: []int{},
		que2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.que2 = append(this.que2, x)
}

func (this *MyQueue) Pop() int {
	if len(this.que1) == 0 {
		for len(this.que2) > 0 {
			this.que1 = append(this.que1, this.que2[len(this.que2)-1])
			this.que2 = this.que2[:len(this.que2)-1]
		}
	}
	x := this.que1[len(this.que1)-1]
	this.que1 = this.que1[:len(this.que1)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.que1) == 0 {
		for len(this.que2) > 0 {
			this.que1 = append(this.que1, this.que2[len(this.que2)-1])
			this.que2 = this.que2[:len(this.que2)-1]
		}
	}
	return this.que1[len(this.que1)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.que1) == 0 && len(this.que2) == 0
}
