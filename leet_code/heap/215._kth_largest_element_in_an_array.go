package main

import (
	"fmt"
)

type Heap struct {
	arr []int
}

func newHeap() Heap {
	return Heap{}
}

// time - O(logn), mem - O(1)
func (this *Heap) push(v int) {
	//добавляем в конец
	this.arr = append(this.arr, v)
	//делаем для него просеивание вверх
	this.shiftUp(len(this.arr) - 1)
}

// time - O(1), mem - O(1)
func (this *Heap) top() int {
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
	if this.arr[parentIdx] > this.arr[i] {
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

/*
	если k вснгда очень малеьнкое - оптимизированнее куча размером k (а не n)
	если k всегда очень большое - оптимизированнее куча размером n

	САМОЕ ОПТИМАЛЬНОЕ РЕШЕНИЕ - QUICK SELECT
	k = 10
	   < pivot >
	100 | 100 | 100
	ищем справа итд
	O(n) - лучший
	O(n^2) - худший
*/
// time - O(k +(n-k)logk), mem - O(n)
func findKthLargest(nums []int, k int) int {
	//с помощью кучи:
	//1) строим MIN!!! кучу
	h := newHeap()
	//2) пушим в кучу первые k элементов
	n := 0
	for n < k {
		h.push(nums[n])
		n++
	}
	//3) сохраняем в куче только максимальные элементы. наверху - самый маленький, поэтоум сравниваем с ним
	for n < len(nums) {
		top := h.top()
		el := nums[n]
		if el > top {
			h.pop()
			h.push(el)
		}
		n++
	}

	//3) получаем элемент
	return h.top()
}

func main() {
	//1 2 2 3 3 4 5 5 6
	n := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(n, k))
}
