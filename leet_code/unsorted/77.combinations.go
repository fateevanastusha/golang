package main

import "fmt"

// func combine(n int, k int) [][]int {

// 	c := n
// 	i := 1
// 	for i <= k {
// 		c = c * (n - i)
// 		i++
// 	}

// 	z := 0
// 	res := [][]int{}

// 	fmt.Println(c/n, n)
// 	//4
// 	for z < n {
// 		x := 0
// 		//6
// 		for x < (c / n) {
// 			arr := []int{z + 1}
// 			res = append(res, arr)
// 			x++
// 		}
// 		z++
// 	}

// 	return res
// }

// n=4, k=3
func combine(n int, k int) [][]int {

	res := [][]int{}

	var rec func(level int, curr []int, start int)
	rec = func(level int, curr []int, start int) {
		if level == k {
			comb := make([]int, k)
			copy(comb, curr)
			res = append(res, comb)
			return
		}
		need := k - level
		for i := start; i <= n-(need-1); i++ {
			rec(level+1, append(curr, i), i+1)
		}
	}

	rec(0, []int{}, 1)
	return res
}

func main() {

	//колво символов
	n := 4
	//длина массива
	k := 3

	fmt.Println(combine(n, k))
}
