package main

import "fmt"

func good(v, t int) bool {
	return v <= t
}

// time - O(logn), mem - O(1)
func search(nums []int, t int) int {
	l, r := 0, len(nums)
	for r-l > 1 {
		m := (l + r) / 2
		if good(nums[m], t) {
			l = m
		} else {
			r = m
		}
	}
	if nums[l] == t {
		return l
	}
	return -1
}

func main() {
	n := []int{-1, 0, 3, 4, 5, 9, 12}
	t := 9
	fmt.Println(search(n, t))
}
