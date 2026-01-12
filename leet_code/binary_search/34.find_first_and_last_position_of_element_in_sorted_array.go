package main

import "fmt"

// time - O(logn), mem - O(1)
func searchRange(nums []int, target int) []int {
	l, r := -1, len(nums)-1
	//find left
	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] < target {
			l = m
		} else {
			r = m
		}
	}
	left := r
	if left != -1 && nums[left] != target {
		return []int{-1, -1}
	}
	l, r = l, len(nums)

	// find right
	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	right := l

	return []int{left, right}
}

func main() {
	n := []int{1, 2, 3, 4, 5}
	t := 7
	fmt.Println(searchRange(n, t))
}
