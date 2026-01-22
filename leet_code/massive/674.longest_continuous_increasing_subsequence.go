package main

import (
	"fmt"
	"math"
)

func findLengthOfLCIS(nums []int) int {
	//time: O(n), mem: O(1)
	var currL, maxL int
	nums = append(nums, math.MaxInt)
	for i, curr := range nums[:len(nums)-1] {
		var next = nums[i+1]
		//если возрастает - поднимаем текущую длину
		if next > curr {
			currL++
		} else {
			//наткнулись на перегиб. обнуляем текущую длину, выставляем максимальную
			if currL+1 > maxL {
				maxL = currL + 1
			}
			currL = 0
		}

	}
	if currL > maxL {
		maxL = currL
	}

	return maxL
}

func main() {
	n := []int{3, 0, 6, 2, 4, 7, 0, 0}
	fmt.Println(findLengthOfLCIS(n))
}
