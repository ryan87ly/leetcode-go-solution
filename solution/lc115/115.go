package lc115

func numDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	dp := make([][]int, len(t))
	for i := 0; i < len(t); i++ {
		dp[i] = make([]int, len(s))
	}
	if s[0] == t[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	for i := 1; i < len(s); i++ {
		if s[i] == t[0] {
			dp[0][i] = dp[0][i-1] + 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if j < i {
				dp[i][j] = 0
			} else {
				if s[j] == t[i] {
					dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[len(t)-1][len(s)-1]
}
