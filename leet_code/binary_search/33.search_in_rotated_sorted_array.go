package main

import "fmt"

// time - O(logn), mem - O(1)
func search(nums []int, target int) int {

	l, r := 0, len(nums)

	//find pivot
	for r-l > 1 {
		m := (l + r) / 2
		if nums[m] > nums[0] {
			l = m
		} else {
			r = m
		}
	}

	offset := l + 1
	l, r = 0, len(nums)

	for r-l > 1 {
		m := (l + r) / 2
		//получаем индекс так, как он должен был быть если бы не сместили массив
		i := (m + offset) % len(nums)
		if nums[i] <= target {
			l = m
		} else {
			r = m
		}
	}

	result := (l + offset) % len(nums)

	if nums[result] != target {
		return -1
	}
	return result
}

func main() {
	m := []int{4, 5, 6, 7, 0, 1, 2, 3}
	t := 1
	fmt.Println(search(m, t))
}
