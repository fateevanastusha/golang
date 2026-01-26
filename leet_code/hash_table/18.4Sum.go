package main

import (
	"fmt"
	"sort"
)

// time - O(n^2), mem - O(n)
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	l := len(nums)
	if l < 4 {
		return res
	}
	sort.Ints(nums)
	//берем первое число и для него перебираем остальное
	for ai, a := range nums[:l-3] {
		//пропускаем дубликаты
		if ai > 0 && a == nums[ai-1] {
			continue
		}
		//берем второе число и для него перебираем остальное
		for bi := ai + 1; bi <= l-3; bi++ {
			b := nums[bi]
			//пропускаем дубликаты
			if bi > ai+1 && b == nums[bi-1] {
				continue
			}
			//классический способ найти сумму
			l, r := bi+1, l-1
			for l < r {
				s := a + b + nums[l] + nums[r]
				if s == target {
					res = append(res, []int{a, b, nums[l], nums[r]})
					l++
					//пропускаем дубликаты
					for l < r && nums[l-1] == nums[l] {
						l++
					}
				}
				if s < target {
					l++
				}
				if s > target {
					r--
				}
			}
		}

	}

	return res
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	t := 0
	fmt.Println(fourSum(nums, t))
}
