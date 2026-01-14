package main

import (
	"fmt"
	"sort"
)

// time - O(n + logn), mem - O(1)
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	a := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	b := nums[0] * nums[1] * nums[len(nums)-1]
	return max(a, b)
}

func main() {
	n := []int{-100, -2, -3, 1}
	fmt.Println(maximumProduct(n))
}
