package main

import (
	"fmt"
	"slices"
	"unicode"
)

func Pop[T comparable](arr *[]T) T {
	a := *arr
	element := a[len(a)-1]
	*arr = a[:len(a)-1]
	return element
}

func minRemoveToMakeValid(str string) string {

	v := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	r := []rune(str)
	stack := []int{}
	result := []rune(str)
	deleted := 0
	for i, s := range r {
		if unicode.IsLetter(s) {
			continue
		}
		//open
		if _, ok := v[s]; ok {
			//push to stack
			stack = append(stack, i)
		} else {
			//close
			if len(stack) == 0 {
				//delete
				result = slices.Delete(result, i-deleted, i-deleted+1)
				deleted++

			} else {
				last := Pop(&stack)
				if v[r[last]] != s {
					//delete
					result = slices.Delete(result, i-deleted, i-deleted+1)
					deleted++

				}
			}

		}
	}
	if len(stack) != 0 {
		for _, v := range stack {
			result = slices.Delete(result, v-deleted, v-deleted+1)
			deleted++
		}
	}

	return string(result)
	// stack := []int{}

	// v := map[rune]rune{
	// 	'[': ']',
	// 	'(': ')',
	// 	'{': '}',
	// }

	// r := []rune(s)
	// res := []rune(s)
	// deleted := 0
	// for i, s := range r {
	// 	//skip letters
	// 	if unicode.IsLetter(s) {
	// 		continue
	// 	}

	// 	if _, ok := v[s]; ok {
	// 		//open
	// 		stack = append(stack, i)
	// 	} else {
	// 		//close

	// 		//no open for close
	// 		if len(stack) == 0 {
	// 			//delete this from string
	// 			res = slices.Delete(res, i-deleted, i-deleted+1)
	// 			deleted++
	// 			continue
	// 		}

	// 		//last open not for this close
	// 		last := Pop(&stack)
	// 		if v[r[last]] != s {
	// 			//delete this from string
	// 			res = slices.Delete(res, i-deleted, i-deleted+1)
	// 			deleted++
	// 			continue
	// 		}
	// 	}

	// }
	// if len(stack) != 0 {
	// 	for _, i := range stack {
	// 		res = slices.Delete(res, i-deleted, i-deleted+1)
	// 		deleted++
	// 	}

	// }

	// return string(res)
}

func main() {
	// s := "lee(t(c)o)de)"
	s := "))(("
	fmt.Println(minRemoveToMakeValid(s))
}
