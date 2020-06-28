package lc448

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		cur := abs(nums[i]) - 1
		if nums[cur] > 0 {
			nums[cur] = -nums[cur]
		}
	}
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
