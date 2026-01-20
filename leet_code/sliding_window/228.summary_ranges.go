package main

import "fmt"

// time - O(n), mem - O(n)
func summaryRanges(nums []int) []string {
	l, r := 0, 0
	res := []string{}
	for l < len(nums) {

		for r+1 < len(nums) && nums[r]+1 == nums[r+1] {
			r++
		}
		if r != l {
			res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r]))
		} else {
			res = append(res, fmt.Sprintf("%d", nums[r]))
		}

		l = r + 1
		r = r + 1
	}
	return res
}

func main() {
	n := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(n))
}
