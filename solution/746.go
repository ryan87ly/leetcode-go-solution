package main

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

func minCostClimbingStairs(cost []int) int {
	steps := len(cost)
	completedCost := append([]int{0}, cost...)
	dp := make([]int, steps+2)
	dp[0], dp[1] = 0, 0
	for i := 2; i < steps+2; i++ {
		dp[i] = min(dp[i-2]+completedCost[i-2], dp[i-1]+completedCost[i-1])
	}
	return dp[steps+1]
}
