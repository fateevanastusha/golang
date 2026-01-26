package main

import (
	"fmt"
	"unicode"
)

func solve(s string) string {
	result := []rune{}
	for _, v := range s {
		if unicode.IsSpace(v) {
			result = append(result, []rune("@40")...)
		}

		result = append(result, v)

	}
	return string(result)
}

func main() {
	//spaces to @40
	s := "Coding Ninjas Is A Coding Platform \nHello World"
	fmt.Println(solve(s))
}
