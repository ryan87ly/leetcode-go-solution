package lc343

func integerBreak(n int) int {
	nums := make([]int, 7)
	nums[2] = 1
	nums[3] = 2
	nums[4] = 4
	nums[5] = 6
	nums[6] = 9
	result := 1
	i := n
	for ; i > 6; i -= 3 {
		result *= 3
	}
	return result * nums[i]
}
