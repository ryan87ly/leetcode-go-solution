package lc264

import "math"

func min(nums ...int) int {
	curMin := math.MaxInt32
	for _, num := range nums {
		if curMin > num {
			curMin = num
		}
	}
	return curMin
}

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for t2, t3, t5, i := 0, 0, 0, 1; i < n; i++ {
		curMin := min(dp[t2]*2, dp[t3]*3, dp[t5]*5)
		if curMin == dp[t2]*2 {
			t2++
		}
		if curMin == dp[t3]*3 {
			t3++
		}
		if curMin == dp[t5]*5 {
			t5++
		}
		dp[i] = curMin
	}
	return dp[n-1]
}
