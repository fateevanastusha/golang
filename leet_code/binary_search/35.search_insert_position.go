package main

import (
	"fmt"
)

// 0 ms
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	nLowMax := 0

	for left <= right {
		mid := (left + right) / 2
		n := nums[mid]
		if n == target {
			return mid
		}
		if n < target {
			if mid > nLowMax {
				nLowMax = mid
			}
			left = mid + 1
		}
		if n > target {
			right = mid - 1
		}
	}

	if target < nums[nLowMax] {
		return nLowMax
	} else {
		return nLowMax + 1
	}
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0

	fmt.Println(searchInsert(nums, target))
}
