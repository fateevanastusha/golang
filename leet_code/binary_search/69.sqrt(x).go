package main

import "fmt"

// time - O(logn), mem - O(1)
func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	l, r := 1, x
	for r-l > 1 {
		m := (l + r) / 2
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

	fmt.Println(mySqrt(2))

}
