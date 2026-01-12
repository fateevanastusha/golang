package main

import (
	"fmt"
)

func Pop[T comparable](arr *[]T) T {
	a := *arr
	element := a[len(a)-1]
	*arr = a[:len(a)-1]
	return element
}

// time - O(n), mem - O(1)
func isValid(str string) bool {

	stack := []rune{}

	v := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	if len(str)%2 != 0 {
		return false
	}

	for _, s := range []rune(str) {
		if _, ok := v[s]; ok {
			// s открывающая
			stack = append(stack, s)
		} else {
			// s закрывающая
			if len(stack) == 0 {
				break
			}
			last := Pop(&stack)
			if v[last] != s {
				break
			}
		}
	}
	return len(stack) == 0
}

func main() {
	s := "{{{{}]}}}"
	fmt.Println(isValid(s))
}
