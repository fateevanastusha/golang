package main

import "fmt"

func getValue(s []rune, i int) rune {
	if i >= len(s) {
		return '!'
	}
	return s[i]
}

// time - O(max(n,m)) - n и m - длины строк, mem - O(1)
func solve(s, t string) bool {
	sR, tR := []rune(s), []rune(t)
	steps, p1, p2 := 0, 0, 0
	for p1 < len(sR) || p2 < len(tR) {
		v1, v2 := getValue(sR, p1), getValue(tR, p2)
		n, m := len(sR)-p1, len(tR)-p2
		if v1 != v2 {
			if n > m {
				p1++
			}
			if m > n {
				p2++
			}
			if m == n {
				p1++
				p2++
			}
			steps++
		} else {
			p1++
			p2++
		}
	}
	return steps == 1
}

func main() {
	s := "aaaaaaa"
	t := "a"
	fmt.Println(solve(s, t))
}
