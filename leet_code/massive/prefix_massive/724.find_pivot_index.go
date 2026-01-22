package main

import (
	"fmt"
)

// time: O(n), mem: O(1)
// префиксный массив
func pivotIndex(nums []int) int {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	prefixSum := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		//слева равно права - префиксная сумма равна общей (общая-текущее значение-префиксная сумма)
		if sum-prefixSum-v == prefixSum {
			return i
		}
		prefixSum += v
	}

	return -1

}

func main() {
	n := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(n))
}
