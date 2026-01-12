package main

import "fmt"

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

func printLevels(levels map[*TreeNode]int) {
	fmt.Println("Levels:")
	for node, level := range levels {
		fmt.Printf("node = %d, level = %d\n", node.Val, level)
	}
}

// 0 ms
func rightSideView(root *TreeNode) []int {
	res := []int{}

	var rec func(node *TreeNode, level int)
	rec = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level >= len(res) {
			res = append(res, node.Val)
		}
		rec(node.Right, level+1)
		rec(node.Left, level+1)
	}
	rec(root, 0)

	// stack := []*TreeNode{root}
	// levels := map[*TreeNode]int{root: 0}
	// for len(stack) != 0 {
	// 	node := Pop(&stack)
	// 	nodeLevel := levels[node]
	// 	if nodeLevel >= len(res) {
	// 		res = append(res, node.Val)
	// 	}
	// 	if node.Left != nil {
	// 		stack = append(stack, node.Left)
	// 		levels[node.Left] = levels[node] + 1
	// 	}
	// 	if node.Right != nil {
	// 		stack = append(stack, node.Right)
	// 		levels[node.Right] = levels[node] + 1
	// 	}
	// }

	return res
}

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   12,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:  8,
				Left: nil,
				Right: &TreeNode{
					Val:   26,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	fmt.Println(rightSideView(root))
}
