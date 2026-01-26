package main

import (
	"fmt"
	"sort"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// time - O(logn*m), mem - O(1)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

	cnt := 0
	//отсортировали второй массив
	sort.Ints(arr2)
	//перебираем значения из первого массива
	for _, n := range arr1 {
		//ответ в l
		l, r := 0, len(arr2)
		f := true
		for r-l > 1 {
			m := (l + r) / 2
			if absInt(arr2[m]-n) <= d {
				f = false
				break
			} else {
				if arr2[m] > n {
					r = m
				} else {
					l = m
				}
			}
		}
		//проверяем валидный ли ответ
		if absInt(arr2[l]-n) <= d {
			f = false
		}
		if f == true {
			cnt++
		}
	}
	return cnt
}

func main() {
	arr1 := []int{8}
	arr2 := []int{-2, 0, 1, 8, 9, 10}
	d := 2
	fmt.Println(findTheDistanceValue(arr1, arr2, d))
}
