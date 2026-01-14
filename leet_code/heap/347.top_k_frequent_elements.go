package main

import "fmt"

type Heap struct {
	arr [][]int
}

func newHeap() Heap {
	return Heap{}
}

// time - O(logn), mem - O(1)
func (this *Heap) push(v []int) {
	//добавляем в конец
	this.arr = append(this.arr, v)
	//делаем для него просеивание вверх
	this.shiftUp(len(this.arr) - 1)
}

// time - O(1), mem - O(1)
func (this *Heap) top() []int {
	return this.arr[0]
}

// time - O(logn), mem - O(1)
func (this *Heap) pop() {
	//меняем местами
	this.arr[0], this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1], this.arr[0]
	//удаляем последний узел
	this.arr = this.arr[0 : len(this.arr)-1]
	//делаем просеивание вниз для первого
	this.shiftDown(0)
}

// time - O(logn), mem - O(1)
func (this *Heap) shiftUp(i int) {
	if i == 0 {
		return
	}
	parentIdx := (i - 1) / 2
	if this.arr[parentIdx][0] > this.arr[i][0] {
		this.arr[parentIdx], this.arr[i] = this.arr[i], this.arr[parentIdx]
		this.shiftUp(parentIdx)
	}
}

// time - O(logn), mem - O(1)
func (this *Heap) shiftDown(i int) {
	leftChildIdx, rightChildIdx := i*2+1, i*2+2
	if leftChildIdx >= len(this.arr) {
		return
	}
	//имеем как минимум левого ребенка
	nextIdx, nextMinVal := leftChildIdx, this.arr[leftChildIdx]
	//если есть правый и он меньше - обновляем
	if rightChildIdx < len(this.arr) && nextMinVal[0] > this.arr[rightChildIdx][0] {
		nextIdx, nextMinVal = rightChildIdx, this.arr[rightChildIdx]
	}
	currentVal := this.arr[i]
	//меняем только если текущее значение больше нижнего
	if nextMinVal[0] < currentVal[0] {
		this.arr[nextIdx], this.arr[i] = currentVal, nextMinVal
		this.shiftDown(nextIdx)
	}

}

func topKFrequent(nums []int, k int) []int {
	//make map
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	//make min-heap with size k
	h := newHeap()
	for key, v := range m {
		//fill heap with k el
		if len(h.arr) < k {
			h.push([]int{v, key})
		} else {
			top := h.top()
			if v > top[0] {
				h.pop()
				h.push([]int{v, key})
			}
		}

	}

	res := []int{}
	for _, v := range h.arr {
		res = append(res, v[1])
	}

	return res

}

func main() {
	n := []int{1, 2, 1, 2, 1, 2, 2, 2, 3, 1, 3, 2, 4, 4, 4}
	k := 2
	fmt.Println(topKFrequent(n, k))
}
