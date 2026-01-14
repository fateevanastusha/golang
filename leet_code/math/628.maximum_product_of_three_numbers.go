package main

import (
	"fmt"
	"sort"
)

// time - O(n + logn), mem - O(1)
func maximumProduct(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		ai, aj := nums[i], nums[j]
		if ai < 0 {
			ai = -ai
		}
		if aj < 0 {
			aj = -aj
		}
		return ai < aj
	})
	r := 1
	for _, v := range nums[len(nums)-3:] {
		r *= v
	}

	return r
}

func main() {
	n := []int{-100, -2, -3, 1}
	fmt.Println(maximumProduct(n))
}
