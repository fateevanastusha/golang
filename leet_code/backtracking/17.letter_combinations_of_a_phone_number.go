package main

import "fmt"

/*
какие сообщения можно набрать на кнопочном телефоне?

	2 - abc
	3 - def
	4 - ghi
	5 - jkl
	6 - mno
	7 - pqrs
	8 - tuv
	9 - wxyz
*/

// func Pop(arr *[]rune) {
// 	a := *arr
// 	*arr = a[:len(a)-1]
// }

// func addLetter(s *[]rune, letter rune) {
// 	*s = append(*s, letter)
// }

// time - O(4**n), mem - O(n*4**n)
func letterCombinations(digits string) []string {

	//рекурсивно
	// res := []string{}
	// if len(digits) == 0 {
	// 	return res
	// }
	// phone := map[rune][]rune{
	// 	'2': []rune("abc"),
	// 	'3': []rune("def"),
	// 	'4': []rune("ghi"),
	// 	'5': []rune("jkl"),
	// 	'6': []rune("mno"),
	// 	'7': []rune("pqrs"),
	// 	'8': []rune("tuv"),
	// 	'9': []rune("wxyz"),
	// }

	// var rec func(idx int, currCombination []rune)
	// rec = func(idx int, currCombination []rune) {
	// 	if idx == len(digits) {
	// 		res = append(res, string(currCombination))
	// 		return
	// 	}
	// 	digit := digits[idx]
	// 	for _, letter := range phone[rune(digit)] {
	// 		//добавили новую букву
	// 		addLetter(&currCombination, letter)
	// 		//перебрали все комбинации для этой буквы
	// 		rec(idx+1, currCombination)
	// 		//удалили ее
	// 		Pop(&currCombination)
	// 	}
	// }

	// rec(0, []rune{})
	// return res

	//периодами
	if len(digits) == 0 {
		return []string{}
	}
	phone := map[rune][]rune{
		'2': []rune("abc"),
		'3': []rune("def"),
		'4': []rune("ghi"),
		'5': []rune("jkl"),
		'6': []rune("mno"),
		'7': []rune("pqrs"),
		'8': []rune("tuv"),
		'9': []rune("wxyz"),
	}

	allComb := 1
	for _, n := range digits {
		allComb *= len(phone[n])
	}
	period := allComb
	res := make([]string, allComb)

	//для каждой цифры пробегаемся allComb раз
	for _, number := range []rune(digits) {
		//длина периода
		period = period / len(phone[number])
		//сколько раз повторяется период
		periodRepeate := allComb / period / len(phone[number])
		//period*periodRepeate*len(phone[number]) = allComb!
		i := 0
		for a := 0; a < periodRepeate; a++ {
			for _, letter := range phone[number] {
				for b := 0; b < period; b++ {
					res[i] = fmt.Sprint(res[i], string(letter))
					i++
				}
			}
		}

	}

	return res

}

func main() {
	d := "723"
	fmt.Println(letterCombinations(d))
}
