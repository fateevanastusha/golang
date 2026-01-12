package main

import "fmt"

// time: O(n), mem: O(n)
func maxPower(s string) int {
	maxL, currL := 0, 0
	b := append([]rune(s), []rune("|")...)
	for i := 0; i <= len(b)-2; i++ {
		currL++
		curr, next := b[i], b[i+1]
		if curr != next {
			if currL > maxL {
				maxL = currL
			}
			currL = 0
		}
	}

	return maxL
}

func main() {
	s := "dddeggggggg"
	fmt.Println(maxPower(s))
}
