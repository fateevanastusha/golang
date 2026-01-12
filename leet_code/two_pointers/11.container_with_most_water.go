package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	res := 0

	l, r := 0, len(height)-1
	for r > l {
		left, right := height[l], height[r]
		curr := (r - l) * min(left, right)

		if curr > res {
			res = curr
		}
		if left <= right {
			l++
		}
		if right <= left {
			r--
		}

	}

	return res
}

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(h))
}
