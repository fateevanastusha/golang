package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	max1, max2 := math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt

	for _, n := range nums {
		if n > max1 {
			max2, max1 = max1, n
		} else if n > max2 {
			max2 = n
		}

		if n < min1 {
			min2, min1 = min1, n
		} else if n < min2 {
			min2 = n
		}
	}

	return max((max1-1)*(max2-1), (min1-1)*(min2-1))
}

func main() {
	n := []int{3, 4, 5, 2}
	fmt.Println(maxProduct(n))

}
