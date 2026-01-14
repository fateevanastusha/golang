package main

import (
	"fmt"
)

// time - O(n), mem - O(n)
func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	for _, v := range nums {
		counter[v]++
	}

	//index - frequency, value - numbers with this frequency
	freq := make([][]int, len(nums)+1)

	for key, frequency := range counter {
		freq[frequency] = append(freq[frequency], key)
	}
	result := []int{}
	for i := len(freq) - 1; i >= 0; i-- {
		for _, num := range freq[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}
	return result

}

func main() {
	n := []int{1, 2, 1, 2, 1, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 3, 2, 4, 4, 4, 4}
	k := 2
	fmt.Println(topKFrequent(n, k))
}
