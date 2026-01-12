package main

import (
	"fmt"
	"sort"
)

// time - O(n^2 + nlogn), mem - O(n)
func threeSum(nums []int) [][]int {

	res := [][]int{}
	sort.Ints(nums)

	for i, x := range nums[:len(nums)-2] {
		if i > 0 && x == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + x
			if sum == 0 {
				res = append(res, []int{x, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
			if sum < 0 {
				l++
			}
			if sum > 0 {
				r--
			}
		}
	}

	return res

	// res := [][]int{}
	// duplicates := map[string]bool{}

	// for i, x := range nums {
	// 	seen := map[int]bool{}
	// 	for _, y := range nums[i+1:] {
	// 		t := 0 - (x + y)
	// 		if _, ok := seen[t]; ok {
	// 			v := []int{x, y, t}
	// 			sort.Ints(v)
	// 			dup := fmt.Sprintf("%v", v)
	// 			if _, ok := duplicates[dup]; !ok {
	// 				res = append(res, v)
	// 				duplicates[dup] = true
	// 			}
	// 		}
	// 		seen[y] = true
	// 	}
	// }
	// return res
}

func main() {
	n := []int{0, 0, 0, 0}
	fmt.Println(threeSum(n))
}
