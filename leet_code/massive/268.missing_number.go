package main

import "fmt"

// time = O(n), mem = O(1)
func missingNumber(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	l := len(nums)
	s := l * (l + 1) / 2
	if s >= sum {
		return s - sum
	}

	/*
		формула арифметической прогрессии:
		первый член = a1
		разность = d
		количество членов = n
		сумма n-членов
		Sn = (n * (a1 + an)) / 2
	*/

	return l
}

func main() {
	n := []int{1}
	fmt.Println(missingNumber(n))
}
