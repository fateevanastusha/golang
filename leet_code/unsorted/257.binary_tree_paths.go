package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0 ms
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var rec func(node *TreeNode, curr string)
	rec = func(node *TreeNode, curr string) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, curr+strconv.Itoa(node.Val))
			return
		}
		rec(node.Left, curr+strconv.Itoa(node.Val)+"->")
		rec(node.Right, curr+strconv.Itoa(node.Val)+"->")
		return
	}

	rec(root, "")
	return res
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(binaryTreePaths(tree))
}
