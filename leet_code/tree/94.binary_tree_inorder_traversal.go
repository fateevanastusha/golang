package main

import "fmt"

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

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var rec func(node *TreeNode)
	rec = func(node *TreeNode) {
		if node == nil {
			return
		}
		rec(node.Left)
		res = append(res, node.Val)
		rec(node.Right)
	}

	rec(root)

	// stack := []*TreeNode{}
	// curr := root
	// for curr != nil || len(stack) > 0 {

	// 	for curr != nil {
	// 		stack = append(stack, curr)
	// 		curr = curr.Left
	// 	}

	// 	curr = Pop(&stack)
	// 	res = append(res, curr.Val)
	// 	curr = curr.Right

	// }

	return res
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

	fmt.Println(inorderTraversal(tree))
}
