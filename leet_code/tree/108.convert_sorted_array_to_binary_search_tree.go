package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//{1, 2, 3, 4, 5, 6, 7, 8, 9}

func sortedArrayToBST(nums []int) *TreeNode {
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := (l + r) / 2
		return &TreeNode{
			Val:   nums[m],
			Left:  build(l, m-1),
			Right: build(m+1, r),
		}
	}
	return build(0, len(nums)-1)
}

//[1,2,3,4,5]
//[1,2,3,4,5,6]

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	node := sortedArrayToBST(nums)
	b, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
