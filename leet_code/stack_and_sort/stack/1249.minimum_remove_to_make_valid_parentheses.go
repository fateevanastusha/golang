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
	//в стэке храним не скобку, а индекс этой скобки. чтобы знать что потом удалять
	stack := []int{}
	result := []rune(str)
	deleted := 0
	for i, s := range r {
		//пропускаем НЕскобки
		if unicode.IsLetter(s) {
			continue
		}
		//открывающая скобка
		if _, ok := v[s]; ok {
			//push to stack
			stack = append(stack, i)
		} else {
			//закрывающая скобка
			//в стэке нет открывающей вообще или последняя открывающая не для нее
			if len(stack) == 0 || v[r[Pop(&stack)]] != s {
				//delete
				//удаляем этот символ, deleted нужна для gap (как указатель сколько мы удалили, чтобы не ошибиться в индексах)
				result = slices.Delete(result, i-deleted, i-deleted+1)
				deleted++

			}

		}
	}

	//удаляем все, что в стэке осталось (открывающие без пары)
	if len(stack) != 0 {
		for _, v := range stack {
			result = slices.Delete(result, v-deleted, v-deleted+1)
			deleted++
		}
	}

	return string(result)
}

func main() {
	// s := "lee(t(c)o)de)"
	s := "))(("
	fmt.Println(minRemoveToMakeValid(s))
}
