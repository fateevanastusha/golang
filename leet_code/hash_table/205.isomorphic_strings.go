package main

import (
	"fmt"
)

//time: O(n), mem: O(n)
/*

	присваиваем значение каждой букве. затем сравниваем значение у первой и второй, они
	должны быть равны.
*/
func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m1, m2 := map[rune]int{}, map[rune]int{}

	s1, t2 := []rune(s), []rune(t)
	for i := 0; i < len(s1); i++ {
		v1, v2 := s1[i], t2[i]
		if _, ok := m1[v1]; !ok {
			m1[v1] = len(m1) + 1
		}

		if _, ok := m2[v2]; !ok {
			m2[v2] = len(m2) + 1
		}

		if m1[v1] != m2[v2] {
			return false
		}

	}

	return true

}
func main() {

	s := "bbbaaaba"
	t := "aaabbbba"

	fmt.Println(isIsomorphic(s, t))

}
