package main

import (
	"fmt"
	"unicode"
)

// time - O(n), mem - O(n)
func reverseWords(s string) string {

	// w := strings.Fields(s)
	// slices.Reverse(w)
	// return strings.Join(w, " ")
	stack := []string{}
	currWord := []rune{}

	for _, v := range []rune(s + " ") {
		if unicode.IsSpace(v) {
			//stop word
			if len(currWord) > 0 {
				stack = append(stack, string(currWord))
				currWord = []rune{}
			}
		} else {
			currWord = append(currWord, v)
		}
		fmt.Println(string(currWord))
	}

	result := ""

	for i := len(stack) - 1; i >= 0; i-- {
		result += stack[i]
		if i != 0 {
			result += " "
		}
	}
	return result
}

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}
