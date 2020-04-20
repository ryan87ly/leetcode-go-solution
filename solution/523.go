package main

import (
	"fmt"
	"math"
)

func mod(num, module int) int {
	if module == 0 && num == 0 {
		return 0
	} else if module == 0 {
		return num
	}
	return int(math.Mod(float64(num), float64(module)))
}

func checkSubarraySum(nums []int, k int) bool {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = -1
	}
	for i := 0; i < length; i++ {
		curNum := nums[i]
		for j := 0; j <= i; j++ {
			isFirst := (dp[j] == -1)
			var preNum int
			if isFirst {
				preNum = 0
			} else {
				preNum = dp[j]
			}

			mod := (curNum + preNum) % k
			dp[j] = mod
			if k == 0 && !isFirst && preNum == 0 && dp[j] == 0 {
				return true
			} else if dp[j] == 0 && !isFirst {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 0))
}
