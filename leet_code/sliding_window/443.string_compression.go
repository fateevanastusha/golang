package main

import "fmt"

// time - O(n), mem - O(1)
func compress(chars []byte) int {
	var cnt int
	l, r, w := 0, 0, 0
	for l < len(chars) {
		curr := 1
		for r+1 < len(chars) && chars[r] == chars[r+1] {
			r++
			curr++
		}
		var iter string
		if curr == 1 {
			iter = fmt.Sprintf("%c", chars[r])
		} else {
			iter = fmt.Sprintf("%c%d", chars[r], curr)
		}
		cnt += len(iter)
		for _, s := range iter {
			chars[w] = byte(s)
			w++
		}
		l = r + 1
		r = r + 1

	}
	return cnt
}

func main() {
	c := []byte{'a', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}
	fmt.Println(compress(c))
	fmt.Println(string(c))
}
