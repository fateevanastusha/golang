package main

import "fmt"

/*
	поиск по ответу. ищем валидный ответ.
*/
// time - O(logn), mem - O(1)
func isPerfectSquare(num int) bool {
	l, r := 1, num
	for (r - l) > 1 {
		m := (r + l) / 2
		s := m * m
		if s <= num {
			l = m
		} else {
			r = m
		}
	}

	return l*l == num
}

func main() {
	fmt.Println(isPerfectSquare(14))
}
