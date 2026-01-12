package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time: O(n), mem: O(n)
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var rec func(node *TreeNode, level int)

	rec = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)

		rec(node.Left, level+1)
		rec(node.Right, level+1)
	}

	rec(root, 0)

	// levels := [][]*TreeNode{{root}}
	// for level := 0; level < len(levels); level++ {
	// 	//collect values from previos level
	// 	prevLevel := levels[level]
	// 	newLevel := []*TreeNode{}
	// 	if len(res)-1 != level {
	// 		res = append(res, []int{})
	// 	}
	// 	for item := 0; item < len(prevLevel); item++ {
	// 		//add item from prev level
	// 		node := prevLevel[item]
	// 		res[level] = append(res[level], node.Val)
	// 		if node.Left != nil {
	// 			newLevel = append(newLevel, node.Left)
	// 		}
	// 		if node.Right != nil {
	// 			newLevel = append(newLevel, node.Right)
	// 		}
	// 	}
	// 	if len(newLevel) > 0 {
	// 		levels = append(levels, newLevel)
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
	fmt.Println(levelOrder(tree))
}
