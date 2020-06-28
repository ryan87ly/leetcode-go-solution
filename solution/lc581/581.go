package lc581

import "sort"

func findUnsortedSubarray(nums []int) int {
	copied := make([]int, len(nums))
	copy(copied, nums)
	sort.Ints(copied)
	identical := 0
	for i := 0; i < len(nums) && copied[i] == nums[i]; i++ {
		identical++
	}
	if identical == len(nums) {
		return 0
	}
	for i := len(nums) - 1; i >= 0 && copied[i] == nums[i]; i-- {
		identical++
	}
	return len(nums) - identical
}
