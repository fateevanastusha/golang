package main

import "fmt"

// time - O(n), mem - O(1)
func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[rune]int, 26)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]--
	}
	for _, cnt := range m {
		if cnt != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "aacc"
	t := "ccac"

	fmt.Println(isAnagram(s, t))
}
