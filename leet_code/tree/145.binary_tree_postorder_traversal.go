package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pop[T any](arr *[]T) T {
	if len(*arr) == 0 {
		var zero T
		return zero
	}
	lastElem := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return lastElem
}

// 0 ms
func postorderTraversal(root *TreeNode) []int {
	res := []int{}

	var rec func(node *TreeNode)
	rec = func(node *TreeNode) {
		if node == nil {
			return
		}

		rec(node.Left)
		rec(node.Right)
		res = append(res, node.Val)
	}

	rec(root)

	// curr := root
	// var visited *TreeNode
	// stack := []*TreeNode{}
	// for curr != nil || len(stack) > 0 {

	// 	//сначала спустимся на самую левую
	// 	for curr != nil {
	// 		stack = append(stack, curr)
	// 		curr = curr.Left
	// 	}

	// 	last := stack[len(stack)-1]
	// 	if last.Right != nil && visited != last.Right {
	// 		curr = last.Right
	// 		visited = last.Right
	// 	} else {
	// 		visited = pop(&stack)
	// 		res = append(res, visited.Val)
	// 	}
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

	fmt.Println(postorderTraversal(tree))
}
