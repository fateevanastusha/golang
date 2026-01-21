package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// рассмотрим почему это работает:
// 1 - если списки одинаковой длины и не имеет пересечений, то получим None
// 2 - если списки одинаковой длины и имеет пересечение, то в каждый раз проверяем
//     пересекаются ли списки и гарантированно получим первую ноду, где пересеклись
// 3 - списки не равной длины и не пересекаются -> получим None т к p1 пройдет в сумме
//     len(headA) + len(headB) и p2 пройдет столько же и оба будут показывать на None
// 4 - списки не равной длины и пересекаются. Тут мы тоже получим, что p1 и p2 пройдут
//     одинаковое число шагов, т к оба идет сначала по headA, а потом headB, значит
//     они будут иметь одинаковый отступ и будут выровнены и придут одновременно
// Note: словами может быть непонятно, поэтому нарисуйте 4 этих случая на бумаге
// и посмотрите как это работает

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB

	// ходим пока не пересечемся!
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}

	return p1
}

func main() {
	common := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: common}}

	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: common,
			},
		},
	}
	fmt.Println(getIntersectionNode(headA, headB))
}
