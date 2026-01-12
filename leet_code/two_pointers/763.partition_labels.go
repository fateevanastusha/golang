package main

import (
	"fmt"
)

func partitionLabels(s string) []int {
	seen := map[byte]int{}
	for i, v := range []byte(s) {
		seen[v] = i
	}
	res := []int{}
	start, end := 0, seen[s[0]]
	for i := 0; i < len(s); i++ {
		last := seen[s[i]]
		if last > end {
			end = last

		}
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

func main() {
	//ababcbaca defegde hijhklij
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}
