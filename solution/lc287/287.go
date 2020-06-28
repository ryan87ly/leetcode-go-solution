package lc287

func nextNIdx(startIdx int, nums []int, n int) int {
	curIdx := startIdx
	for i := 0; i < n; i++ {
		curIdx = nums[curIdx]
	}
	return curIdx
}

func findDuplicate(nums []int) int {
	slowIdx, fastIdx := 0, 0
	for {
		slowIdx = nextNIdx(slowIdx, nums, 1)
		fastIdx = nextNIdx(fastIdx, nums, 2)
		if nums[slowIdx] == nums[fastIdx] {
			break
		}
	}
	headIdx := 0
	for nums[headIdx] != nums[slowIdx] {
		headIdx = nextNIdx(headIdx, nums, 1)
		slowIdx = nextNIdx(slowIdx, nums, 1)
	}
	return nums[headIdx]

}
