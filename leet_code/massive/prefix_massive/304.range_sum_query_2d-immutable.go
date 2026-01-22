package main

import "fmt"

type NumMatrix struct {
	m [][]int
}

/*

	индексация префиксного массива

	0 0 0 0
	0 1 1 1
	0 1 1 1
	0 1 1 1

	0 0 0 0
	0 1 2 3
	0 2 4 6
	0 3 6 9

	(получается что каждый квадрат это сумма всех чисел)

	0 0 0 | 0
	0 1 2 | 3
	0 2 4 | 6
	------
	0 3 6   9

	складываем верхнее и левое число, и вычитаем по диагонали (так как мы его два раза добавили) + добавляем
	собственное число этой клетки
	4 = 2 + 2 - 1 + 1
	2 + 2 - левое и верхнее
	- 1 - по диагонали
	+ 1 - собственное число клетки



	чтобы найти сумму элементов в квадрате делаем так:

	выделяем квадрат, вычетаем из него лишние квадраты и добавляем по диагонали
	(так как вычли два раза). это и будет результат
	0 0   0 0
	0 1   2 3
			----
	0 2 | 4 6
	0 3 | 6 9

	например, чтобы найти для нижнего правого квадрата 2x2 сумму элементов -
	вычитаем левый (3), вычитаем верхний (3), добавляем диагональный (1)
	9 - 3 - 3 + 1 = 4
*/

// time: 0(n), mem: O(1)
func ConstructorM(matrix [][]int) NumMatrix {

	//append row
	row := []int{}
	for _, _ = range matrix[0] {
		row = append(row, 0)
	}
	matrix = append([][]int{row}, matrix...)
	//append col
	for i, _ := range matrix {
		matrix[i] = append([]int{0}, matrix[i]...)
	}

	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			matrix[row][col] = matrix[row-1][col] + matrix[row][col-1] - matrix[row-1][col-1] + matrix[row][col]
		}
	}

	return NumMatrix{
		m: matrix,
	}
}

// time: 0(1), mem: O(1)
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1, col1, row2, col2 = row1+1, col1+1, row2+1, col2+1
	return this.m[row2][col2] - this.m[row2][col1-1] - this.m[row1-1][col2] + this.m[row1-1][col1-1]
}

func main() {
	m := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	// m := [][]int{
	// 	{1, 1, 1},
	// 	{1, 1, 1},
	// 	{1, 1, 1},
	// }
	nm := ConstructorM(m)

	r1, c1 := 1, 1
	r2, c2 := 3, 3

	fmt.Println(nm.SumRegion(r1, c1, r2, c2))
}
