package algobase

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	INF int = 0x7fffffff
)

// Generate a binary tree from int slice
// Time: O(n)
// Space: O(n)
func BuildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := make(map[int]*TreeNode)
	// Firtly, genereate all tree nodes
	for i, v := range nums {
		if v == INF {
			m[i] = nil
		} else {
			m[i] = &TreeNode{v, nil, nil}
		}
	}
	// Then, set relationship according to order of elements
	for i, v := range nums {
		if v != INF {
			m[i].Left = m[2*i+1]
			m[i].Right = m[2*i+2]
		}
	}
	return m[0]
}
