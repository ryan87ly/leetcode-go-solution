package lc78

func copyList(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

func backtrack(curr []int, nums []int, curIdx, targetLen int) [][]int {
	if len(curr) == targetLen {
		return [][]int{curr}
	}
	result := make([][]int, 0)
	for i := curIdx; i < len(nums); i++ {
		num := nums[i]
		nextCur := append(copyList(curr), num)
		result = append(result, backtrack(nextCur, nums, i+1, targetLen)...)
	}
	return result
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	emptySlice := make([]int, 0)
	for i := 0; i <= len(nums); i++ {
		result = append(result, backtrack(copyList(emptySlice), nums, 0, i)...)
	}
	return result
}
