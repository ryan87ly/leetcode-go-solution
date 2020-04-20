package lc121

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxValue, lowestBefore := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		cur := prices[i] - lowestBefore
		maxValue = max(maxValue, cur)
		lowestBefore = min(lowestBefore, prices[i])
	}
	return maxValue
}
