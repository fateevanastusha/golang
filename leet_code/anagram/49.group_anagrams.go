package main

import (
	"fmt"
	"slices"
)

// time - O(n+mlogm) (n-количество символов), mem - O(n)
func groupAnagrams(strs []string) [][]string {

	res := [][]string{}
	m := map[string][]string{}

	for _, s := range strs {
		r := []rune(s)
		slices.Sort(r)
		m[string(r)] = append(m[string(r)], s)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	r := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(r))
}
