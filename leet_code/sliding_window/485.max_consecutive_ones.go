package main

import "fmt"

// time - O(n), mem - O(1)
func findMaxConsecutiveOnes(nums []int) int {
	l, r := 0, 0
	var max int
	for l < len(nums) {

		for r+1 < len(nums) && nums[r] == nums[r+1] {
			r++
		}
		if nums[r] == 1 {
			c := r - l + 1
			if c > max {
				max = c
			}
		}

		l = r + 1
		r = r + 1

	}

	return max
}

func main() {
	n := []int{0}
	fmt.Println(findMaxConsecutiveOnes(n))
}
