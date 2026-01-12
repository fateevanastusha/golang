package main

import (
	"fmt"
	"strings"
	"unicode"
)

func longestNiceSubstring(s string) string {
	allLetters := make(map[rune]int)
	for _, l := range []rune(s) {
		allLetters[l] = 1
	}
	swap := func(s rune) rune {
		if unicode.IsLower(s) {
			return unicode.ToUpper(s)
		}
		return unicode.ToLower(s)
	}
	forReplace := make(map[rune]int)
	for key, _ := range allLetters {
		pair := swap(key)
		fmt.Println(string(key), string(pair))
		if _, ok := allLetters[pair]; !ok {
			forReplace[key] = 1
		}
	}
	for key, _ := range forReplace {
		s = strings.ReplaceAll(s, string(key), ".")
	}
	splitted := strings.Split(s, ".")
	max := ""
	for _, str := range splitted {

		if len(str) > len(max) {
			max = str
		}
	}
	return max
}

func main() {
	s := "YazaAay"
	fmt.Println(longestNiceSubstring(s))
}
