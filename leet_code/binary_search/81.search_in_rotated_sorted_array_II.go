package main

import "fmt"

// time - O(n)
func search(nums []int, target int) bool {
	// for _, v := range nums {
	// 	if v == target {
	// 		return true
	// 	}
	// }
	// return false

	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return true
		}

		// не можем определить сторону — сдвигаем границу
		//[2 2 3 4 5 1 2 2] - пример (слева и справа одинаковые числа)
		if nums[l] == nums[m] && nums[m] == nums[r] {
			l++
			r--
		} else if nums[l] <= nums[m] {
			//левый меньше таргерта и таргет меньше миддл
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
			//левый больше таргета
		} else {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return false

}

func main() {
	n := []int{2, 5, 6, 0, 0, 1, 2}
	t := 2
	fmt.Println(search(n, t))
}
