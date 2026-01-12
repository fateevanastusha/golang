package main

import (
	"fmt"
)

// time - O(n), mem - O(n)
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i := len(nums) - 1
	l, r := 0, len(nums)-1
	for r >= l {
		left, right := nums[l]*nums[l], nums[r]*nums[r]
		if left > right {
			result[i] = left
			l++
		} else {
			result[i] = right
			r--
		}
		i--

	}
	return result
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
