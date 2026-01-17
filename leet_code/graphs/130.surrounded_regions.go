package main

import "fmt"

// time - O(m*n), mem - O(m*n)
/*
	проходимся по всем границам. если находим остров - то красим все, что к нему примыкает.
	после того, как все острова, примыкающие к границам, окрашены - проходимся по остальным
	островам (минуя примыкающие) и делаем их X.
*/
func solve(board [][]byte) {
	used := make([][]bool, len(board))
	for i, v := range board {
		used[i] = make([]bool, len(v))
	}
	var dfs func(a, b int)
	dfs = func(a, b int) {
		if (a < len(board) && a >= 0 && b < len(board[0]) && b >= 0 && board[a][b] != 'X' && used[a][b] == false) == false {
			return
		}
		used[a][b] = true
		steps := [][]int{
			//down
			{a + 1, b},
			//left
			{a, b - 1},
			//up
			{a - 1, b},
			//right
			{a, b + 1},
		}
		for _, step := range steps {
			dfs(step[0], step[1])
		}

	}

	//check boards to find islands that примыкают to borders
	for a := 0; a < len(board); a++ {
		if a == 0 || a == len(board)-1 {
			for b := 0; b < len(board[a]); b++ {
				if board[a][b] != 'X' {
					dfs(a, b)
				}
			}
		} else {
			for b := 0; b < len(board[a]); b += len(board[a]) - 1 {
				if board[a][b] != 'X' {
					dfs(a, b)
				}
			}
		}
	}

	//iter inside boards
	for a := 1; a < len(board)-1; a++ {
		for b := 1; b < len(board[a])-1; b++ {
			if board[a][b] == 'X' || used[a][b] == true {
				continue
			}
			board[a][b] = 'X'
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(board)

	fmt.Println(board)
}
