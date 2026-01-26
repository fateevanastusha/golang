package main

import "fmt"

func check(s string, l, r int) bool {
	for r > l {
		if s[r] != s[l] {
			return false
		}
		r--
		l++
	}

	return true
}

/*
можно удалить максимум один символ (0 или 1)
*/
//time - O(n), mem - O(1)
func validPalindromeII(s string) bool {
	l, r := 0, len(s)-1
	for r > l {
		if s[r] != s[l] {
			//проверяем, валидные ли последовательности без этих символов.
			return check(s, l+1, r) || check(s, l, r-1)
		}

		r--
		l++
	}
	return true
}

func main() {
	s := "deeee"
	fmt.Println(validPalindromeII(s))
}
