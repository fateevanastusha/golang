package main

import "fmt"

type MyQueue struct {
	storage []int
}

func ConstructorA() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.storage = append(this.storage, x)
}

func (this *MyQueue) Pop() int {

	if len(this.storage) == 0 {
		return -1
	}
	v := this.storage[0]
	this.storage = this.storage[1:]
	return v

}

func (this *MyQueue) Peek() int {
	if len(this.storage) == 0 {
		return -1
	}
	v := this.storage[0]
	return v
}

func (this *MyQueue) Empty() bool {
	return len(this.storage) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	q := ConstructorA()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())  //1
	fmt.Println(q.Pop())   //1
	fmt.Println(q.Empty()) //false
}
