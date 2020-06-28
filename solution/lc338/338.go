package lc338

func countBits(num int) []int {
	result := make([]int, num + 1)
	result[0] = 0
	for i:=1; i <= num; i++ {
		result[i] = result[i >> 1] + i % 2
	}
	return result
}