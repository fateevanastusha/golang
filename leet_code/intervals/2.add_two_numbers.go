package main

import "fmt"

func addTwoNumbers(arr1, arr2 []int) []int {
	l := max(len(arr1), len(arr2))
	res := make([]int, l+1)
	f1, f2 := len(arr1)-1, len(arr2)-1

	//идем с конца
	for f1 >= 0 || f2 >= 0 {

		index := max(f1, f2)
		//сумма чисел
		var r int
		if f1 >= 0 {
			r += arr1[f1]
		}
		if f2 >= 0 {
			r += arr2[f2]
		}

		//кладем в конец
		res[index+1] += r % 10
		//если большие десяти - то в следующий разряд
		res[index] += r / 10

		f1--
		f2--
	}
	return res
}

func main() {
	arr1, arr2 := []int{1, 2, 7}, []int{4, 5, 6, 4}
	fmt.Println(addTwoNumbers(arr1, arr2))
}
