package main

import "fmt"

// time - O(logn), mem - O(1)
func findMin(nums []int) int {
	//4,5,6,7,0,1,2
	/*
		здесь надо просто найти offset

		4 5 6 7 | 0 1 2
		BAD      GOOD

		good всегда меньше или равны последнему элементу массива
	*/
	l, r := -1, len(nums)-1

	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] <= nums[len(nums)-1] {
			r = m
		} else {
			l = m
		}

	}
	return nums[r]
}

func main() {
	n := []int{11, 12, 13, 17}
	fmt.Println(findMin(n))
}
