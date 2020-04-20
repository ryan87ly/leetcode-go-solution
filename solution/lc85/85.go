package lc85

import "math"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	rows := len(matrix)
	columns := len(matrix[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	totalMax := 0
	for i := 0; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i][j-1] + 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			curMin := math.MaxInt32
			for k := i; k >= 0 && dp[k][j] > 0; k-- {
				curMin = min(curMin, dp[k][j])
				length := i - k + 1
				totalMax = max(totalMax, curMin * length)
			}
		}
	}
	return totalMax
}