package main

import "fmt"

func solve(A string) int {
	var s int
	for _, v := range []rune(A) {
		if v == '(' {
			s++
		} else {
			s--
		}
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
