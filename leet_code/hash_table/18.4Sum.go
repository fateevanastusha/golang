package main

import (
	"fmt"
	"sort"
)

// time - O(n^2), mem - O(n)
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	l := len(nums)
	if l < 4 {
		return res
	}
	sort.Ints(nums)
	for ai, a := range nums[:l-3] {

		if ai > 0 && a == nums[ai-1] {
			continue
		}
		for bi := ai + 1; bi <= l-3; bi++ {
			b := nums[bi]

			if bi > ai+1 && b == nums[bi-1] {
				continue
			}
			l, r := bi+1, l-1
			for l < r {
				s := a + b + nums[l] + nums[r]
				if s == target {
					res = append(res, []int{a, b, nums[l], nums[r]})
					l++
					for l < r && nums[l-1] == nums[l] {
						l++
					}
				}
				if s < target {
					l++
				}
				if s > target {
					r--
				}
			}
		}

	}

	return res
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	t := 0
	fmt.Println(fourSum(nums, t))
}
