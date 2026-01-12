package main

import "fmt"

// time - O(n), mem - O(1)
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for r > l {
		sum := numbers[l] + numbers[r]
		if sum == target {
			break
		}
		if sum < target {
			l++
		}
		if sum > target {
			r--
		}
	}

	return []int{l + 1, r + 1}
}

func main() {
	n := []int{2, 7, 11, 15}
	t := 9

	fmt.Println(twoSumII(n, t))
}
