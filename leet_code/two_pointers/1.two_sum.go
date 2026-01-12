package main

import "fmt"

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		left, right := nums[l], nums[r]
		sum := left + right
		if sum == target {
			return []int{l, r}
		}
		if sum < target {
			l++
		}
		if sum > target {
			r--
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 9))

}
