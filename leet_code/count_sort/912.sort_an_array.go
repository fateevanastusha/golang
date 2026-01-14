package main

import (
	"fmt"
	"math"
	"sort"
)

// time - O(n+klogk), mem - O(k)
func sortArray(nums []int) []int {
	/*
		//time - O(n), mem - O(n)

				# считаем сколько каких чисел в массиве
		        # ключ - число, значение - сколько встретилось

		        # 5 * 10 ** 4 - т к минимальный элемент это
		        # -5 * 10 ** 4 (по условию), то сдвиг будет
		        # +5 * 10 ** 4, чтобы count[0] - показывал
		        # скалько элементов в массиве равных -5 * 10 ** 4
		        offset = 5 * 10 ** 4
		        count = [0 for _ in range(5 * 10 ** 4 * 2 + 1)]
		        for num in nums:
		            count[num + offset] += 1

		        # я создаю новый а вы при желании можете изменить массив nums
		        sorted_nums = []
		        for num in range(len(count)):
		            for _ in range(count[num]):
		                sorted_nums.append(num - offset)
		        return sorted_nums
	*/
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
	n := []int{5, 2, 3, 1}
	fmt.Println(sortArray(n))
}
