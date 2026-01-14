package main

import "fmt"

// min-куча
type Heap struct {
	arr []int
}

func newHeap() Heap {
	return Heap{}
}

func (this *Heap) push(v int) {
	//добавляем в конец
	this.arr = append(this.arr, v)
	//делаем для него просеивание вверх
	this.shiftUp(len(this.arr) - 1)
}
func (this *Heap) top() int {
	return this.arr[len(this.arr)-1]
}

func (this *Heap) pop() {
	//меняем местами
	this.arr[0], this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1], this.arr[0]
	//удаляем последний узел
	this.arr = this.arr[0 : len(this.arr)-1]
	//делаем просеивание вниз для первого
	this.shiftDown(0)
}

func (this *Heap) shiftUp(i int) {
	if i == 0 {
		return
	}
	parentIdx := (i - 1) / 2
	if this.arr[parentIdx] > this.arr[i] {
		this.arr[parentIdx], this.arr[i] = this.arr[i], this.arr[parentIdx]
		this.shiftUp(parentIdx)
	}
}

func (this *Heap) shiftDown(i int) {
	leftChildIdx, rightChildIdx := i*2+1, i*2+2
	if leftChildIdx >= len(this.arr) {
		return
	}
	//имеем как минимум левого ребенка
	nextIdx, nextMinVal := leftChildIdx, this.arr[leftChildIdx]
	//если есть правый и он меньше - обновляем
	if rightChildIdx < len(this.arr) && nextMinVal > this.arr[rightChildIdx] {
		nextIdx, nextMinVal = rightChildIdx, this.arr[rightChildIdx]
	}
	currentVal := this.arr[i]
	//меняем только если текущее значение больше нижнего
	if nextMinVal < currentVal {
		this.arr[nextIdx], this.arr[i] = currentVal, nextMinVal
		this.shiftDown(nextIdx)
	}

}

func main() {
	h := newHeap()
	h.push(3)
	h.push(2)
	h.push(1)
	fmt.Println(h.arr)
	h.pop()
	fmt.Println(h.arr)
}
