package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	p, l := []rune(pattern), (strings.Split(s, " "))
	if len(p) != len(l) {
		return false
	}
	patterns, letters := map[rune]int{}, map[string]int{}

	for i := 0; i < len(p); i++ {
		pv, lv := p[i], l[i]
		if _, ok := patterns[pv]; !ok {
			patterns[pv] = len(patterns) + 1
		}
		if _, ok := letters[lv]; !ok {
			letters[lv] = len(letters) + 1
		}

		if patterns[pv] != letters[lv] {
			return false
		}
	}

	return true
}

func main() {

	pattern := "abba"
	s := "dog cat cat dog"

	fmt.Println(wordPattern(pattern, s))

}
