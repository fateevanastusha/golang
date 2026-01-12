package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Pop[T any](arr *[]T) T {
	if len(*arr) == 0 {
		var zero T
		return zero
	}
	lastElem := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return lastElem
}

func isValidBST(root *TreeNode) bool {

	var rec func(root *TreeNode, min int, max int) bool

	rec = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}

		left := rec(node.Left, min, node.Val)
		right := rec(node.Right, node.Val, max)
		return left && right
	}

	return rec(root, math.MinInt64, math.MaxInt64)

}

func main() {

	// t1 := &TreeNode{
	// 	Val: 5,
	// 	Left: &TreeNode{
	// 		Val:   3,
	// 		Left:  &TreeNode{Val: 2},
	// 		Right: &TreeNode{Val: 4},
	// 	},
	// 	Right: &TreeNode{
	// 		Val:   7,
	// 		Left:  &TreeNode{Val: 6},
	// 		Right: &TreeNode{Val: 8},
	// 	},
	// }

	t2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(isValidBST(t2))
}
