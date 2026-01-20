package main

import (
	"fmt"
	"strings"
)

/*
решение ТАКОЕ ЖЕ! только с проверкой колонки/строки/диагонали
*/
func findPositions(currPosition [][]string, n, queenRow int, usedRows, usedCols, usedLrDiag, usedRlDiag []bool, allPositions *[][]string) {
	//все, пришли, комбинация удачная!
	if n == 0 {

		//добавляем в ответ
		res := []string{}
		for _, v := range currPosition {
			res = append(res, strings.Join(v, ""))
		}

		*allPositions = append(*allPositions, res)
		return
	}
	if usedRows[queenRow] == true {
		return
	}

	for queenCol := range currPosition {
		//проверяем, не занята ли строка/колонка/диагональ другой королевой
		if usedCols[queenCol] || usedRows[queenRow] || usedLrDiag[queenRow-queenCol+len(currPosition)] || usedRlDiag[queenRow+queenCol] {
			continue
		}
		//ставим королеву
		usedCols[queenCol] = true
		usedRows[queenRow] = true
		usedLrDiag[queenRow-queenCol+len(currPosition)] = true
		usedRlDiag[queenRow+queenCol] = true
		currPosition[queenRow][queenCol] = "Q"

		//ищем отвлетвления от этой позиции для следующих строк
		findPositions(currPosition, n-1, queenRow+1, usedRows, usedCols, usedLrDiag, usedRlDiag, allPositions)

		//убираем королеву
		usedCols[queenCol] = false
		usedRows[queenRow] = false
		usedLrDiag[queenRow-queenCol+len(currPosition)] = false
		usedRlDiag[queenRow+queenCol] = false
		currPosition[queenRow][queenCol] = "."
	}
}

func solveNQueens(n int) [][]string {

	if n == 1 {
		return [][]string{{"Q"}}
	}

	/*
		в каждой диагонали, строке и колонке МОЖЕТ БЫТЬ ТОЛЬКО ОДНА КОРОЛЕВА.
		поэтому мы нашли крутой способ запомнить, занята ли эта диагональ, колонка или строка уже другой королевой.
	*/
	currPosition := make([][]string, n)
	for a := 0; a < n; a++ {
		currPosition[a] = make([]string, n)
		for b := 0; b < n; b++ {
			currPosition[a][b] = "."
		}
	}
	//колонки и строки - тут все просто. индекс - номер колонки или строки
	rows, cols := make([]bool, n), make([]bool, n)
	//диагональ слева направо (сверху вниз) и справа налево (снизу вверх). формула поиска номера диагонали выше.
	lrDiag, rlDiag := make([]bool, n*2), make([]bool, n*2)
	var positions [][]string
	findPositions(currPosition, n, 0, rows, cols, lrDiag, rlDiag, &positions)

	return positions
}

func main() {
	n := 2
	fmt.Println(solveNQueens(n))
}
