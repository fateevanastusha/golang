package main

import (
	"fmt"
)

func rotateSubArray(nums []int, i, j int) {
	//j not including
	j--
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	//time: O(1), mem: O(1)
	// last := len(nums)
	// n := k % last
	// copy(nums, append(nums[last-n:], nums[0:last-n]...))

	//time: O(n), mem: O(1)
	//1 2 3 4 5 6 7,  k = 3
	l := len(nums)
	n := k % l
	//> 7 6 5 4 3 2 1
	rotateSubArray(nums, 0, l)
	//> 5 6 7 4 3 2 1
	rotateSubArray(nums, 0, n)
	//> 5 6 7 1 2 3 4
	rotateSubArray(nums, n, l)

}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7}
	k := 2

	rotate(n, k)
	fmt.Println(n)
}
