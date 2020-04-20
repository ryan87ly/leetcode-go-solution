package lc213

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	length := len(nums)
	dp := make([]int, length)

	//start with 0
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < length-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	startWithZero := dp[length-2]

	//start with 1
	dp[0] = 0
	dp[1] = nums[1]
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	startWithOne := dp[length-1]

	return max(startWithZero, startWithOne)
}
