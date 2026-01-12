package main

import (
	"fmt"
	"sort"
)

// time - O(n*logn), mem - O(n)
func topKFrequent(nums []int, k int) []int {

	/*
		// 1. Подсчет частот
		count := make(map[int]int)
		for _, num := range nums {
			count[num]++
		}

		// 2. Создаем buckets (bucket sort)
		// Индекс = частота, значение = числа с такой частотой
		frequencyList := make([][]int, len(nums)+1)
		for num, freq := range count {
			frequencyList[freq] = append(frequencyList[freq], num)
		}

		// 3. Собираем результат с конца (самые частые)
		result := make([]int, 0, k)
		for i := len(frequencyList) - 1; i >= 0; i-- {
			if len(frequencyList[i]) > 0 {
				for _, num := range frequencyList[i] {
					result = append(result, num)
					k--
					if k == 0 {
						return result
					}
				}
			}
		}

		return result
	*/
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	freq := make([][2]int, 0, len(m))
	for key, value := range m {
		freq = append(freq, [2]int{key, value})
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i][1] > freq[j][1]
	})

	n := make([]int, 0, k)
	for i := 0; i < k; i++ {
		n = append(n, freq[i][0])
	}

	return n[0:k]
}

func main() {
	n := []int{1, 1, 1, 2, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 8, 8, 8, 8, 8, 9}
	k := 2
	fmt.Println(topKFrequent(n, k))
}
