package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Height struct {
	Left  int
	Right int
}

//0 ms

func diameterOfBinaryTree(root *TreeNode) int {
	m := map[*TreeNode]*Height{}
	var height func(node *TreeNode, currHeight int) int
	height = func(node *TreeNode, currHeight int) int {
		if node == nil {
			return 0
		}
		left := height(node.Left, currHeight+1)
		right := height(node.Right, currHeight+1)

		m[node] = &Height{
			Left:  left,
			Right: right,
		}

		if left > right {
			return 1 + left
		} else {
			return 1 + right
		}
	}
	height(root, 0)

	max := 0

	for _, h := range m {
		if d := h.Left + h.Right; d > max {
			max = d
		}
	}

	return max
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	fmt.Println(diameterOfBinaryTree(tree))
}
