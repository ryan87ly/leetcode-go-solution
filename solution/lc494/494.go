package lc494

func findTargetSumWays(nums []int, S int) int {
	numLen := len(nums)
	dp := make([]map[int]int, numLen+1)
	for i := 0; i <= numLen; i++ {
		dp[i] = make(map[int]int)
	}
	dp[0][0] = 1
	for i := 0; i < numLen; i++ {
		for sum, count := range dp[i] {
			plusSum := sum + nums[i]
			dp[i+1][plusSum] = dp[i+1][plusSum] + count
			minusSum := sum - nums[i]
			dp[i+1][minusSum] = dp[i+1][minusSum] + count
		}
	}
	return dp[numLen][S]
}
