package main

import (
	"math"
	"sort"
)

func sortArray(nums []int) []int {
	min, max := math.MaxInt, math.MinInt
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	keys := []int{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	z := 0
	for _, key := range keys {
		for i := 0; i < m[key]; i++ {
			nums[z] = key
			z++
		}
	}

	return nums
}

func main() {

}
