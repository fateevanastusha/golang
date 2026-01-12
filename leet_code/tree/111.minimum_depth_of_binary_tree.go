package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Pop[T comparable](arr *[]T) T {
	a := *arr
	element := a[len(a)-1]
	*arr = a[:len(a)-1]
	return element
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 && right != 0 {
		return right + 1
	}
	if right == 0 && left != 0 {
		return left + 1
	}
	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}

func main() {

	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:  4,
					Left: nil,
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		Right: nil,
	}

	fmt.Println(minDepth(tree))
}
