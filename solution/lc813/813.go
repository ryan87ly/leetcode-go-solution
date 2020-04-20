package lc813

import (
	"math"
)

func avg(A []int, startIndex, endIndex int) float64 {
	sum := 0
	for i := startIndex; i < endIndex; i++ {
		sum += A[i]
	}
	return float64(sum) / float64(endIndex-startIndex)
}

func largestSumOfAverages(A []int, K int) float64 {
	dp := make([][]float64, K)
	for i := 0; i < K; i++ {
		dp[i] = make([]float64, len(A))
	}

	//initilzation
	for i := 0; i < K; i++ {
		dp[i][0] = float64(A[0])
	}
	for i := 1; i < len(A); i++ {
		dp[0][i] = avg(A, 0, i+1)
	}

	for i := 1; i < K; i++ {
		for j := 1; j < len(A); j++ {
			max := 0.0
			for m := 0; m < j; m++ {
				value := dp[i-1][m] + avg(A, m+1, j+1)
				max = math.Max(max, value)
			}
			dp[i][j] = max
		}
	}
	return dp[K-1][len(A)-1]
}
