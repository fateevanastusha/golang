package main

import "fmt"

type MinStack struct {
	storage    []int
	minStorage []int
	top        int
}

func ConstructorM() MinStack {
	return MinStack{
		top: -1,
	}
}

func (this *MinStack) Push(val int) {
	this.storage = append(this.storage, val)
	if len(this.minStorage) == 0 || val <= this.minStorage[len(this.minStorage)-1] {
		this.minStorage = append(this.minStorage, val)
	}
	this.top++

}

func (this *MinStack) Pop() {
	if len(this.storage) == 0 {
		return
	}
	if this.storage[this.top] == this.minStorage[len(this.minStorage)-1] {
		this.minStorage = this.minStorage[0 : len(this.minStorage)-1]
	}
	this.storage = this.storage[0:this.top]
	this.top--

}

func (this *MinStack) Top() int {
	if this.top == -1 {
		return -1
	}
	return this.storage[this.top]
}

func (this *MinStack) GetMin() int {
	if len(this.minStorage) == 0 {
		return -1
	}
	return this.minStorage[len(this.minStorage)-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	ops := []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin", "pop", "push", "push", "push", "getMin", "pop", "top", "getMin", "pop", "getMin", "pop"}
	args := [][]int{{}, {0}, {1}, {0}, {}, {}, {}, {}, {}, {}, {-2}, {-1}, {-2}, {}, {}, {}, {}, {}, {}, {}}

	var ms MinStack

	for i, op := range ops {
		switch op {
		case "MinStack":
			ms = ConstructorM()
		case "push":
			ms.Push(args[i][0])
		case "pop":
			ms.Pop()
		case "top":
			(ms.Top())
		case "getMin":
			(ms.GetMin())
		}
		fmt.Println(ms.storage, op)

	}
}
