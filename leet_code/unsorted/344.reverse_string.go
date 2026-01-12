package main

import "fmt"

func reverseString(s []byte) {

	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

}

func main() {

	s := "Hannah"
	b := []byte(s)
	reverseString(b)
	fmt.Println(string(b))

}
