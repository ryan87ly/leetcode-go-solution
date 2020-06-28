package lc104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func md(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	return max(md(cur.Left), md(cur.Right)) + 1
}

func maxDepth(root *TreeNode) int {
	return md(root)
}
