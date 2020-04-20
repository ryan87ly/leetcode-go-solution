package lc312

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxCoins(nums []int) int {
	n := len(nums)
	appendedLength := n + 2
	appendedNums := make([]int, appendedLength)
	appendedNums[0] = 1
	appendedNums[n+1] = 1
	for i := 1; i <= n; i++ {
		appendedNums[i] = nums[i-1]
	}

	dp := make([][]int, appendedLength)
	for i := 0; i < appendedLength; i++ {
		dp[i] = make([]int, appendedLength)
	}

	for gap := 2; gap < appendedLength; gap++ {
		for i := 0; i < appendedLength-gap; i++ {
			j := i + gap
			curMax := 0
			for m := i + 1; m < j; m++ {
				cur := dp[i][m] + appendedNums[i]*appendedNums[m]*appendedNums[j] + dp[m][j]
				curMax = max(cur, curMax)
			}
			dp[i][j] = curMax
		}
	}
	return dp[0][n+1]
}
