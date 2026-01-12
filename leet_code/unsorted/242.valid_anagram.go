package main

import (
	"fmt"
	"slices"
)

func isAnagram(s string, t string) bool {
	sB, sT := ([]byte(s)), []byte(t)
	slices.Sort(sB)
	slices.Sort(sT)

	if len(sB) != len(sT) {
		return false
	}
	if len(sB) == 1 {
		return sB[0] == sT[0]
	}
	l, r := 0, len(sB)-1

	for l < r {
		if sB[l] != sT[l] || sB[r] != sT[r] {
			return false
		}
		l++
		r--
	}
	return true
}
func main() {

	s := "hannah"
	t := "ahnnah"
	fmt.Println(isAnagram(s, t))

}
