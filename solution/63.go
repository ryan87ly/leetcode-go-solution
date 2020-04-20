package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])
	dp := make([][]int, height)
	for i := 0; i < height; i++ {
		dp[i] = make([]int, width)
	}

	//initlization
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	for i := 1; i < height; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = 0
		}
	}
	for i := 1; i < width; i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = 0
		}
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}

	return dp[height-1][width-1]
}

// func main() {
// 	fmt.Println(uniquePathsWithObstacles(
// 		[][]int{
// 			[]int{0, 0, 0},
// 			[]int{0, 1, 0},
// 			[]int{0, 0, 0},
// 		},
// 	))
// }
