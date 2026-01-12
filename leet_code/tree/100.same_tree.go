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

// time: O(n), mem: O(h)
func isSameTree(p *TreeNode, q *TreeNode) bool {

	var rec func(p *TreeNode, q *TreeNode) bool
	rec = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}

		return rec(p.Left, q.Left) && rec(p.Right, q.Right)
	}

	return rec(p, q)
	// pStack, qStack := []*TreeNode{p}, []*TreeNode{q}

	// for len(pStack) != 0 || len(qStack) != 0 {
	// 	if len(pStack) != len(qStack) {
	// 		return false
	// 	}
	// 	pNode, qNode := Pop(&pStack), Pop(&qStack)

	// 	if (pNode == nil || qNode == nil) && pNode != qNode {
	// 		return false
	// 	}
	// 	if pNode != nil && qNode != nil && pNode.Val != qNode.Val {
	// 		return false
	// 	}

	// 	if pNode != nil && (pNode.Left != nil || pNode.Right != nil) {
	// 		pStack = append(pStack, pNode.Left)
	// 		pStack = append(pStack, pNode.Right)
	// 	}

	// 	if qNode != nil && (qNode.Right != nil || qNode.Left != nil) {
	// 		qStack = append(qStack, qNode.Left)
	// 		qStack = append(qStack, qNode.Right)
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
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(isSameTree(tree, tree))

}
