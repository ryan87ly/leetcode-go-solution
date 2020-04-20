package lc689

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	m := 3
	subArrayNums := n - k + 1
	subArrays := make([]int, subArrayNums)
	dp := make([][]int, subArrayNums)
	startingIndices := make([][][]int, subArrayNums)

	for i := 0; i < subArrayNums; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			sum += nums[i+j]
		}
		subArrays[i] = sum
	}

	for i := 0; i < subArrayNums; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < subArrayNums; i++ {
		startingIndices[i] = make([][]int, m+1)
	}

	// initialization
	for j := 1; j <= m; j++ {
		dp[0][j] = subArrays[0]
		startingIndices[0][j] = []int{0}
	}

	for i := 1; i < subArrayNums; i++ {
		current := subArrays[i]
		if current > dp[i-1][1] {
			dp[i][1] = current
			startingIndices[i][1] = []int{i}
		} else {
			dp[i][1] = dp[i-1][1]
			startingIndices[i][1] = startingIndices[i-1][1]
		}
	}

	for j := 2; j <= m; j++ {
		for i := 1; i < subArrayNums; i++ {
			if i < k {
				if dp[i-1][j] >= dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
					startingIndices[i][j] = startingIndices[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
					startingIndices[i][j] = startingIndices[i][j-1]
				}
			} else {
				if dp[i-1][j] >= dp[i][j-1] && dp[i-1][j] >= (dp[i-k][j-1]+subArrays[i]) {
					dp[i][j] = dp[i-1][j]
					startingIndices[i][j] = startingIndices[i-1][j]
				} else if dp[i][j-1] >= dp[i-1][j] && dp[i][j-1] >= (dp[i-k][j-1]+subArrays[i]) {
					dp[i][j] = dp[i][j-1]
					startingIndices[i][j] = startingIndices[i][j-1]
				} else {
					dp[i][j] = dp[i-k][j-1] + subArrays[i]
					startingIndices[i][j] = make([]int, len(startingIndices[i-k][j-1]))
					copy(startingIndices[i][j], startingIndices[i-k][j-1])
					startingIndices[i][j] = append(startingIndices[i][j], i)
				}
			}
		}
	}
	return startingIndices[subArrayNums-1][m]
}
