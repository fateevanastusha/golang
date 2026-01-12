package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printMap(m map[*TreeNode][]*TreeNode) {
	for key, vals := range m {
		fmt.Printf("Key(%p) = %d: [", key, key.Val)
		for i, v := range vals {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d", v.Val)
		}
		fmt.Println("]")
	}
}

/*func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	m := map[*TreeNode][]*TreeNode{root: []*TreeNode{}}
	var lastMatch *TreeNode
	var rec func(node *TreeNode, prev *TreeNode)
	rec = func(node *TreeNode, prev *TreeNode) {
		if lastMatch != nil {
			return
		}
		if prev != nil {
			m[node] = append(append(m[node], m[prev]...), prev)
		}
		pa, qa := m[p], m[q]
		for i := 0; i < len(pa) || i < len(qa); i++ {
			if i >= len(pa) {
				if qa[i] == p {
					lastMatch = p
				}
			} else if i >= len(qa) {
				if pa[i] == q {
					lastMatch = q
				}
			} else {
				if pa[i] == qa[i] {
					lastMatch = pa[i]
				}
			}
		}
		if node.Left != nil {
			rec(node.Left, node)
		}
		if node.Right != nil {
			rec(node.Right, node)
		}
	}
	rec(root, nil)
	return lastMatch
}
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func main() {

	var root = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right))
}
