package main

import (
	"fmt"
)

func romanToInt(s string) int {
	str := []byte(s)
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var sum int

	for i := 0; i <= len(s)-1; i++ {
		current := m[string(str[i])]
		if (i+1 < len(s)) && m[string(str[i+1])] > current {
			sum += m[string(str[i+1])] - current
			i++
		} else {
			sum += current
		}

	}

	return sum

}

func main() {
	fmt.Println(romanToInt("III"))

}
