package lc53

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]

	for i, previousSum := 1, maxSum; i < len(nums); i++ {
		if previousSum < 0 {
			previousSum = nums[i]
		} else {
			previousSum += nums[i]
		}
		maxSum = max(maxSum, previousSum)
	}

	return maxSum
}
