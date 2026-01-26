package main

import "fmt"

/*
	бинарный поиск по ответу
	то есть мы рассматриваем ответы, и смотрим - является ли оно ответом.
*/
// time - O(logn), mem - O(1)
func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	l, r := 1, x
	for r-l > 1 {
		//m -  предполагаемый ответ
		m := (l + r) / 2
		//получаем квадрат, и смотрим, меньше он или больше целевого.
		s := m * m
		if s <= x {
			l = m
		} else {
			r = m
		}
	}

	return l

}

func main() {

	fmt.Println(mySqrt(9))

}
