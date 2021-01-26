package week1

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func getTargetCopy(original, cloned, target *TreeNode) *TreeNode {
	cloned = cloneNode(original)
	return cloned
}

func cloneNode(original *TreeNode) *TreeNode {
	if original == nil {
		return nil
	}
	target := new(TreeNode)
	target.Val = original.Val
	target.Left = cloneNode(original.Left)
	target.Right = cloneNode(original.Right)
	return target
}
