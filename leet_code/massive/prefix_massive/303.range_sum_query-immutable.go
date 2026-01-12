package main

import "fmt"

type NumArray struct {
	storage []int
}

// time: 0(n), mem: O(n)
func Constructor(nums []int) NumArray {

	p := []int{0}
	for i, v := range nums {
		p = append(p, v+p[i])
	}

	return NumArray{
		storage: p,
	}
}

// time: 0(1), mem: O(1)
func (this *NumArray) SumRange(left int, right int) int {
	return this.storage[right+1] - this.storage[left]
}

func main() {

	n := []int{-2, 0, 3, -5, 2, -1}
	num := Constructor(n)
	fmt.Println(num.SumRange(0, 2)) //1
	fmt.Println(num.SumRange(2, 5)) // -1
	fmt.Println(num.SumRange(0, 5)) //-3

}
