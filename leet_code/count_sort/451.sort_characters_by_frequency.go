package main

import (
	"fmt"
)

// time - O(n), mem - O(n)
func frequencySort(s string) string {
	count := map[rune]int{}
	for _, v := range []rune(s) {
		count[v]++
	}

	freq := make([][]rune, len(s)+1)

	for key, frequency := range count {
		freq[frequency] = append(freq[frequency], key)
	}

	res := []rune{}
	for i := len(freq) - 1; i >= 0; i-- {
		for _, r := range freq[i] {
			cnt := 0
			for cnt < i {
				res = append(res, r)
				cnt++
			}
		}
	}

	return string(res)

}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}
