package main

import "fmt"

// time - O(logn), mem - O(1)
func findMin(nums []int) int {
	l, r := -1, len(nums)-1
	for r-l > 1 {
		m := (l + r) / 2
		if nums[m] > nums[len(nums)-1] {
			l = m
		} else {
			r = m
		}
	}
	return nums[r]
}

func main() {
	n := []int{4, 5, 6, 7, -1, 0, 1, 2}
	fmt.Println(findMin(n))
}
