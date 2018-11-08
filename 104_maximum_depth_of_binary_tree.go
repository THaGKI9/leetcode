package leetcode

// Source: https://leetcode.com/problems/maximum-depth-of-binary-tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

/*
Test Case:
[3,9,20,null,null,15,7]
*/
