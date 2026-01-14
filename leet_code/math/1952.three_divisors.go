package main

import (
	"fmt"
)

// time - O(n), mem - O(1)
func isThree(n int) bool {
	flag := false
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			if flag {
				return false
			}
			flag = true
		}

	}
	return flag
}

func main() {
	n := 14
	//1 14 2 7
	fmt.Println(isThree(n))
}
