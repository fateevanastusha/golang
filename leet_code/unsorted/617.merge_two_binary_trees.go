package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func heightTree(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	return 1 + heightTree(tree.Left) + heightTree(tree.Right)
}

//0 ms

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	height1 := heightTree(root1)
	height2 := heightTree(root2)
	l := 0

	stack := [][]*TreeNode{{root1}}
	for l < height1 {
		arr := []*TreeNode{}
		previos := stack[l-1]

		for 
		l++
	}

}

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Left: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(mergeTrees(tree1, tree2))
}
