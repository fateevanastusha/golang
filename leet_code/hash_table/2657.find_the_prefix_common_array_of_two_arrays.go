package main

import "fmt"

// time - O(n^2), mem - 0(n)
/*

	нужно найти количество общих чисел!
	[1 4 3 2]
	[3 2 1 4]

	0:
	[1]
	[3]
	= 0

	1:
	[1 4]
	[3 2]
	= 0

	2:
	[1 4 3]
	[3 2 1]
	= 2

	3:
	[1 4 3 2]
	[3 2 1 4]
	= 4

	каждый раз в эти массивы ставим TRUE, когда нашли значение (значение из искомого массива -
	индекс в этих), затем смотрим количество пересечений, где TT на одном индексе стоит

	A = [..., F F F F]
	B = [..., F F F F]

	1)
	A = [..., T F F F]
	B = [..., F F T F]
	= 0

	2)
	A = [..., T F F T]
	B = [..., F T T F]
	= 0

	3)
	A = [..., T F T T]
	B = [..., T T T F]
	= 2

	4)
	A = [..., T T T T]
	B = [..., T T T T]
	= 4



	Из этого вытекает битовое решение тоже
*/
func findThePrefixCommonArray(A []int, B []int) []int {

	//time - O(n^2), mem - O(n)
	l := len(A)
	sA, sB := make([]bool, l+1), make([]bool, l+1)
	r := make([]int, l)
	for i := 0; i < l; i++ {
		a, b := A[i], B[i]
		//ставим true
		sA[a], sB[b] = true, true

		var cnt int

		//смотрим количество пересечений
		for k := 1; k < l+1; k++ {
			if sA[k] && sB[k] {
				cnt++
			}
		}

		r[i] = cnt

	}

	return r

	//time - O(n), mem - O(n)
	// 	l := len(A)
	// r := make([]int, l)
	// var sA, sB uint
	// for i := 0; i < l; i++ {

	// 	//same sA[A[i]] = True, добавляем единицу на нужный знак
	// 	sA = sA | 1<<A[i]
	// 	sB = sB | 1<<B[i]

	// 	var cnt uint
	// 	//get intersection
	// если setA = "101100" (в двоичном предствалении)
	// если setB = "111000" (в двоичном предствалении)
	// то setA & setB это:
	//             "101000"
	// 	n := sA & sB

	// 	//collect cnt of 1 (same as True)
	// 	for n != 0 {
	// 		cnt += n % 2
	// 		n /= 2
	// 	}
	// 	r[i] = int(cnt)

	// }
	// return r
}

func main() {
	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}
	fmt.Println(findThePrefixCommonArray(A, B))
}
