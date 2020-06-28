package lc162

func find(nums []int, startIdx, endIdx int) int {
	if startIdx == endIdx {
		return startIdx
	}
	midIdx := (startIdx + endIdx) / 2
	if midIdx == 0 {
		return 1
	}
	midElem := nums[midIdx]
	leftElem := nums[midIdx-1]
	rightElem := nums[midIdx+1]
	if midElem > leftElem && midElem > rightElem {
		return midIdx
	} else if midElem > leftElem && midElem < rightElem {
		return find(nums, midIdx+1, endIdx)
	} else if midElem < leftElem && midElem > rightElem {
		return find(nums, startIdx, midIdx-1)
	} else {
		return find(nums, startIdx, midIdx-1)
	}
}

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if nums[0] > nums[1] {
		return 0
	} else if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	} else {
		return find(nums, 0, len(nums)-1)
	}

}
