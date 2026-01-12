package main

import "fmt"

type MyStack struct {
	storage []int
}

func ConstructorZ() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.storage = append(this.storage, x)
}

func (this *MyStack) Pop() int {
	l := len(this.storage)
	if l == 0 {
		return -1
	}
	last := this.storage[l-1]
	this.storage = this.storage[0 : l-1]
	return last
}

func (this *MyStack) Top() int {
	l := len(this.storage)
	if l == 0 {
		return -1
	}
	return this.storage[l-1]
}

func (this *MyStack) Empty() bool {
	return len(this.storage) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	o := ConstructorZ()
	o.Push(1)
	o.Push(2)
	fmt.Println(o.Top())
	fmt.Println(o.Pop())
	fmt.Println(o.Empty())
}
