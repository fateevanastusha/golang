package main

import (
	"fmt"
)

/*
1 1 2 2 4

0^1=1
1^1=0
0^2=2
2^2=0
0^4=4

В КАКОМ ПОРЯДКЕ НЕ ВАЖНО!
*/
func singleNumber(nums []int) int {

	// time - O(n), mem - O(1)
	/*
		1 1 2 2 4

		0^1=1
		1^1=0
		0^2=2
		2^2=0
		0^4=4

		В КАКОМ ПОРЯДКЕ НЕ ВАЖНО!

	*/
	// var a int
	// for _, n := range nums {
	// 	a = n ^ a
	// }
	// return a
}

func main() {
	n := []int{-1}
	fmt.Println(singleNumber(n))
}
