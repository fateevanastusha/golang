package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// time - O(n), mem - O(n)
func nextLargerNodes(head *ListNode) []int {
	values := []int{}
	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}

	result := make([]int, len(values))
	stack := []int{0}

	for i := 1; i < len(values); i++ {
		v := values[i]

		for len(stack) > 0 && v > values[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[last] = v
		}

		stack = append(stack, i)
		fmt.Println(stack)

	}

	for len(stack) != 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result[last] = 0
	}

	return result

}

func main() {

	node := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(nextLargerNodes(node))

}
