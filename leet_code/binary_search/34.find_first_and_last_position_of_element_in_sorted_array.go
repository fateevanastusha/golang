package main

import "fmt"

// time - O(logn), mem - O(1)
func searchRange(nums []int, target int) []int {
	/*
		тут мы ищем сначала где число начинается, а потом где число заканчивается


	*/
	/*
		5 | 7,7,8,8,10
		BAD GOOD

		BAD - меньше чем target
		GOOD - больше или равен таргету
	*/
	l, r := -1, len(nums)-1
	//find left
	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m
		}
	}
	//ответ в r
	left := r
	//если начало не равно таргету - то его вообще нет, выходим
	if left != -1 && nums[left] != target {
		return []int{-1, -1}
	}

	/*
		5 7 7 | 8 8 10
		GOOD   BAD

		BAD - больше чем target
		GOOD - меньше или равен таргету

		ответ в левой
	*/
	l, r = l, len(nums)

	// find right
	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	right := l

	return []int{left, right}
}

func main() {
	n := []int{1, 2, 3, 4, 5}
	t := 7
	fmt.Println(searchRange(n, t))
}
