package lc309

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
	dp0 := make([]int, n+1)
	dp1 := make([]int, n+1)

	dp0[0], dp0[1] = 0, 0
	dp1[0], dp1[1] = 0, -prices[0]

	for i := 2; i <= n; i++ {
		dp0[i] = max(dp1[i-1]+prices[i-1], dp0[i-1])
		dp1[i] = max(dp0[i-2]-prices[i-1], dp1[i-1])
	}

	return max(dp0[n], dp1[n])
}
