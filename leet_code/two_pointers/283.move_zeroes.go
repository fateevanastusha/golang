package main

import "fmt"

func moveZeroes(nums []int) {
	p1, p2 := 0, 0
	for p2 < len(nums) {
		if nums[p2] != 0 {
			nums[p1] = nums[p2]
			p1++
		}
		p2++
	}
	for p1 < len(nums) {
		nums[p1] = 0
		p1++
	}

}

func main() {
	n := []int{0, 1, 0, 3, 12}
	moveZeroes(n)
	fmt.Println(n)
}
