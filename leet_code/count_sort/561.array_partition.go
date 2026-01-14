package main

import (
	"fmt"
	"slices"
)

// time - O(n + logn), mem - O(1)
func arrayPairSum(nums []int) int {
	slices.Sort(nums)
	var sum int
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	n := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(n))
}
