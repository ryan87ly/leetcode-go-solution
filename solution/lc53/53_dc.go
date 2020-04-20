package lc53

import "math"

type division struct {
	left, right, sum, max int
}

func maxNum(nums ...int) int {
	curMax := math.MinInt32
	for _, num := range nums {
		if curMax < num {
			curMax = num
		}
	}
	return curMax
}

func divideAndConque(nums []int, start, end int) division {
	n := end - start
	if n == 1 {
		return division{nums[start], nums[start], nums[start], nums[start]}
	}
	midIndex := start + (end-start)/2
	leftDivision, rightDivision := divideAndConque(nums, start, midIndex), divideAndConque(nums, midIndex, end)
	left := maxNum(leftDivision.left, leftDivision.sum+rightDivision.left)
	right := maxNum(rightDivision.right, rightDivision.sum+leftDivision.right)
	sum := leftDivision.sum + rightDivision.sum
	max := maxNum(leftDivision.max, rightDivision.max, leftDivision.right+rightDivision.left)
	return division{left, right, sum, max}
}

func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return divideAndConque(nums, 0, len(nums)).max
}
