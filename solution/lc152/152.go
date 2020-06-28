package lc152

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

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	globalMax, curMin, curMax := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = max(num, curMin*num), min(num, curMax*num)
		} else if num > 0 {
			curMax, curMin = max(num, curMax*num), min(num, curMin*num)
		} else {
			curMin, curMax = 0, 0
		}
		globalMax = max(globalMax, curMax)
	}
	return globalMax
}
