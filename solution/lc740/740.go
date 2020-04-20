package lc740

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int)
	for _, value := range nums {
		m[value]++
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	counts := make([]int, len(keys))
	for i, k := range keys {
		counts[i] = m[k]
	}
	dp := make([]int, len(keys))
	dp[0] = keys[0] * counts[0]
	if len(keys) == 1 {
		return dp[0]
	}
	if (keys[1] - keys[0]) == 1 {
		dp[1] = max(dp[0], keys[1]*counts[1])
	} else {
		dp[1] = dp[0] + keys[1]*counts[1]
	}
	if len(keys) == 2 {
		return dp[1]
	}
	for i := 2; i < len(keys); i++ {
		if (keys[i] - keys[i-1]) == 1 {
			dp[i] = max(dp[i-1], dp[i-2]+keys[i]*counts[i])
		} else {
			dp[i] = dp[i-1] + keys[i]*counts[i]
		}
	}
	return dp[len(keys)-1]
}
