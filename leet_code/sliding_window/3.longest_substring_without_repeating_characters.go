package main

import "fmt"

// time - O(n), mem - O(k), где k - количество уникальных символов в строке (мощность алфавита)
func lengthOfLongestSubstring(s string) int {
	l, r := 0, -1
	str := []rune(s)
	var res int
	/*
		какие буквы находястя внутри текущего окна
	*/
	hash := map[rune]bool{}
	for l < len(s) {
		for r+1 < len(str) && hash[str[r+1]] == false {
			r++
			hash[str[r]] = true
		}
		res = max(r-l+1, res)

		/*
			чтобы рассмотреть все варианты - левый двигаем на один вперед, а предыдущее значение
			удаляем из мапы (так как мы его уже не используем)
		*/
		delete(hash, str[l])
		l++

	}

	return res
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
