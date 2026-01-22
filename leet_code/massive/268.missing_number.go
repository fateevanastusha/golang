package main

import "fmt"

// time = O(n), mem = O(1)
/*
	последовательность с 0 до n, неотсортированная
	ищем сумму текущих членов
	ищем нужную сумму (по длине, с 0 до len(nums))

	вычитаем
	если нужная больше - то не хватило числа в середине
	если нужная равна - то 0 не хватило
	если нужная меньше - значит последнего члена не хватило (len(nums))
*/
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
