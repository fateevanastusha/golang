package main

import "fmt"

// 0 ms
func moveZeroes(nums []int) {
	l, r := 0, 0
	for r <= len(nums)-1 {
		if nums[r] != 0 {
			if l != -1 {
				nums[l], nums[r] = nums[r], nums[l]
				l++
			}
		}
		r++

	}
}

func main() {
	n := []int{0, 1, 0, 3, 12}
	moveZeroes(n)
	fmt.Println(n)

}
