package main

import "fmt"

// здесь у нас всего один тип скобок, поэтому баланса достаточно
// time - O(n), mem - O(1)
func solve(A string) int {
	var s int
	/*
		открывающая - плюс
		закрывающая - минус

		по сути так же как стэк
	*/
	for _, v := range []rune(A) {
		if v == '(' {
			s++
		} else {
			s--
		}

		//значит для закрывающей нет открывающей, выходим
		if s < 0 {
			return 0
		}
	}
	if s == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	A := "(()())"
	fmt.Println(solve(A))
}
