package main

import "fmt"

func pop[T any](arr *[]T) T {
	if len(*arr) == 0 {
		var zero T
		return zero
	}
	lastElem := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return lastElem
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time: 0(n), mem: O(h) (h = высота дерева)
func preorderTraversal(root *TreeNode) []int {
	res := []int{}

	var rec func(node *TreeNode)
	rec = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		rec(node.Left)
		rec(node.Right)
	}
	rec(root)

	return res
	// stack := []*TreeNode{root}

	// for len(stack) != 0 {
	// 	curr := pop(&stack)
	// 	res = append(res, curr.Val)
	// 	if curr.Right != nil {
	// 		stack = append(stack, curr.Right)
	// 	}
	// 	if curr.Left != nil {
	// 		stack = append(stack, curr.Left)
	// 	}
	// }

	// return res
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
	fmt.Println(preorderTraversal(tree))
}
