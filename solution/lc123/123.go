package lc123

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k := 2
	dp0 := make([][]int, n)
	dp1 := make([][]int, n)
	for i := 0; i < n; i++ {
		dp0[i] = make([]int, k+1)
		dp1[i] = make([]int, k+1)
	}

	// initilzation
	for i := 0; i < n; i++ {
		dp0[i][0] = 0
		dp1[i][0] = math.MinInt32
	}

	for i := 0; i <= k; i++ {
		dp0[0][i] = 0
		dp1[0][i] = -prices[0]
	}

	// run
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp0[i][j] = max(dp1[i-1][j]+prices[i], dp0[i-1][j])
			dp1[i][j] = max(dp0[i-1][j-1]-prices[i], dp1[i-1][j])
		}
	}

	return max(dp0[n-1][k], dp1[n-1][k])
}
