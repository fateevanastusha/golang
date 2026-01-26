package main

import "fmt"

// time: O(n), mem: O(n)
func twoSum(nums []int, target int) []int {
	//запоминаем те что видели
	seen := make(map[int]int)
	for i, n := range nums {
		c := target - n
		//смотрим есть ли нужное число для суммы
		if idx, ok := seen[c]; ok {
			return []int{idx, i}
		}
		seen[n] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{22, 7, 11, 15, 6, 4, 3}, 9))

}
