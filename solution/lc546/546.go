package lc546

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func calculate(dp [][][]int, boxes []int, i, j, k int) int {
	if i > j {
		return 0
	} else if dp[i][j][k] != -1 {
		return dp[i][j][k]
	} else if i == j {
		cur := (k + 1) * (k + 1)
		dp[i][j][k] = cur
		return cur
	} else {
		curMax := (k+1)*(k+1) + calculate(dp, boxes, i+1, j, 0)
		for m := i + 1; m <= j; m++ {
			if boxes[i] == boxes[m] {
				cur := calculate(dp, boxes, i+1, m-1, 0) + calculate(dp, boxes, m, j, k+1)
				curMax = max(cur, curMax)
			}
		}
		dp[i][j][k] = curMax
		return curMax
	}
}

func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
			for k := 0; k < n; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	return calculate(dp, boxes, 0, n-1, 0)
}
