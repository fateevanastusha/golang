package main

import (
	"fmt"
	"unicode"
)

func lengthOfLastWord(s string) int {
	var res int
	sr := []rune(s)
	i := len(sr) - 1
	for i >= 0 && unicode.IsSpace(sr[i]) {
		i--
	}
	for i >= 0 && !unicode.IsSpace(sr[i]) {
		res++
		i--
	}
	return res
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))

}
