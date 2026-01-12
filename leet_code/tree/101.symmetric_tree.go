package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
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

// time: 0(n), mem: O(h)
func isSymmetric(root *TreeNode) bool {

	var rec func(left *TreeNode, right *TreeNode) bool
	rec = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		return rec(left.Left, right.Right) && rec(left.Right, right.Left)
	}

	return rec(root.Left, root.Right)

	// left := root.Left
	// right := root.Right

	// stackLeft := []*TreeNode{left}
	// stackRight := []*TreeNode{right}
	// for len(stackLeft) != 0 || len(stackRight) != 0 {
	// 	if len(stackLeft) == 0 || len(stackRight) == 0 {
	// 		return false
	// 	}
	// 	l, r := Pop(&stackLeft), Pop(&stackRight)
	// 	if (l == nil || r == nil) && l != r {
	// 		return false
	// 	}
	// 	if l != nil && r != nil && l.Val != r.Val {
	// 		return false
	// 	}

	// 	if l != nil && (l.Left != nil || l.Right != nil) {
	// 		stackLeft = append(stackLeft, l.Left)
	// 		stackLeft = append(stackLeft, l.Right)
	// 	}

	// 	if r != nil && (r.Right != nil || r.Left != nil) {
	// 		stackRight = append(stackRight, r.Right)
	// 		stackRight = append(stackRight, r.Left)
	// 	}
	// }

	// return true

}

func main() {
	// [5,4,8,11,null,13,4,7,2,null,null,null,1]
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isSymmetric(tree))

}
