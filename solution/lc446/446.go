package lc446

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	dp := make([]map[int]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int64)
	}

	res := 0
	for right := 0; right < n; right++ {
		for left := 0; left < right; left++ {
			diff := A[right] - A[left]
			cur := dp[right][diff]
			cur += 1
			dp[right][diff] = cur
			if leftValue, found := dp[left][diff]; found {
				dp[right][diff] = cur + leftValue
				res += int(leftValue)
			}
		}
	}
	return res
}
