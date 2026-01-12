package main

import "fmt"

func repeatedNumber(m []int) []int {
	s_expected, s_actual, sq_expected, sq_actual := 0, 0, 0, 0
	for i, v := range m {
		s_actual += v
		sq_actual += v * v
		sq_expected += (i + 1) * (i + 1)
		s_expected += i + 1
	}
	A := s_actual - s_expected
	B := sq_actual - sq_expected
	x, y := (B+A*A)/(2*A)-A, (B+A*A)/(2*A)
	return []int{y, x}
}

func main() {
	ints := []int{3, 1, 2, 5, 3}
	fmt.Println(repeatedNumber(ints))
}
