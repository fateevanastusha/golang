package main

import "fmt"

// time - O(m*n), mem - O(n*m)
func uniquePaths(m int, n int) int {
	m++
	n++
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
	}
	a[1][1] = 1

	for m := 1; m < len(a); m++ {
		for n := 1; n < len(a[0]); n++ {
			if m == 1 && n == 1 {
				continue
			}
			a[m][n] = a[m][n-1] + a[m-1][n]
		}
	}
	return a[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
