package main

import "fmt"

func sortColors(nums []int) {
	l, m, h := 0, 0, len(nums)-1
	for m <= h {
		if nums[m] == 0 {
			nums[l], nums[m] = nums[m], nums[l]
			l++
			m++
		} else if nums[m] == 1 {
			m++
			continue
		} else {
			nums[m], nums[h] = nums[h], nums[m]
			h--
		}

	}
}

func main() {

	n := []int{2, 0, 2, 1, 1, 0}
	sortColors(n)
	fmt.Println(n)
}
