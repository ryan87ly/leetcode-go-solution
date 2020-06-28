package lc437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func backTrack(cur *TreeNode, target int) int {
	if cur == nil {
		return 0
	}
	count := 0
	if cur.Val == target {
		count++
	}
	return count + backTrack(cur.Left, target-cur.Val) + backTrack(cur.Right, target-cur.Val)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return backTrack(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
