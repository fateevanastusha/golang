package main

import (
	"fmt"
	"strings"
)

func Pop(arr *[]string) {
	a := *arr
	*arr = a[:len(a)-1]
}

func Add(s *[]string, l string) {
	*s = append(*s, l)
}

// time - O(2**n), mem - O(2**n)
/*
	по аналогии с 17/47/51 и задачей с одним видом скобок с балансом
*/
func generateParenthesis(n int) []string {

	var res []string
	l := n * 2

	var rec func(idx int, curr []string)

	rec = func(balance int, curr []string) {
		if balance < 0 || balance > l-len(curr) {
			return
		}
		if balance == 0 && len(curr) == l {
			res = append(res, strings.Join(curr, ""))
			return
		}

		for _, p := range []string{"(", ")"} {
			var newBalance = 1
			if p == ")" {
				newBalance = -1
			}
			Add(&curr, p)
			rec(balance+newBalance, curr)
			Pop(&curr)
		}
	}

	rec(0, []string{})

	return res
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))

}
