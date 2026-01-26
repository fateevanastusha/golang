package main

import "fmt"

// time - O(n+m), mem - O(n), где n - самый короткий массив из двух

/*
идем двумя указателями по массиву. если они равны - в результат. если нет - МЕНЬШИЙ
двигаем.
*/
func intersect(A []int, B []int) []int {
	res := []int{}
	a, b := 0, 0
	for a < len(A) && b < len(B) {
		aValue, bValue := A[a], B[b]
		if aValue == bValue {
			res = append(res, aValue)
			a++
			b++
		}
		if aValue > bValue {
			b++
		}
		if bValue > aValue {
			a++
		}
	}
	return res
}

func main() {
	// Тест 1: обычный случай с повторами
	A1 := []int{1, 2, 3, 3, 4, 5, 6}
	B1 := []int{3, 3, 5}
	fmt.Println(intersect(A1, B1)) // [3 3 5]

	// Тест 2: без повторов в одном из массивов
	A2 := []int{1, 2, 3, 3, 4, 5, 6}
	B2 := []int{3, 5}
	fmt.Println(intersect(A2, B2)) // [3 5]

	// Тест 3: пустое пересечение
	A3 := []int{1, 2, 3}
	B3 := []int{4, 5, 6}
	fmt.Println(intersect(A3, B3)) // []

	// Тест 4: идентичные массивы
	A4 := []int{1, 1, 2, 2}
	B4 := []int{1, 1, 2, 2}
	fmt.Println(intersect(A4, B4)) // [1 1 2 2]

	// Тест 5: один массив пустой
	A5 := []int{1, 2, 3}
	B5 := []int{}
	fmt.Println(intersect(A5, B5)) // []

	// Тест 6: разные длины
	A6 := []int{1, 1, 2, 3, 4}
	B6 := []int{1, 1, 1, 2}
	fmt.Println(intersect(A6, B6)) // [1 1 2]
}
