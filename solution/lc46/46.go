package lc46

func delete(is []int, index int) []int {
	firstSlice := is[:index]
	first := make([]int, len(firstSlice))
	copy(first, firstSlice)
	secondSlice := is[index+1:]
	second := make([]int, len(secondSlice))
	copy(second, secondSlice)

	return append(first, second...)
}

func compute(remainingNums []int, cur []int) [][]int {
	if len(remainingNums) == 0 {
		return [][]int{cur}
	}
	result := make([][]int, 0)
	for i, num := range remainingNums {
		nextRemaining := delete(remainingNums, i)
		nextSet := make([]int, len(cur))
		copy(nextSet, cur)
		nextSet = append(nextSet, num)
		result = append(result, compute(nextRemaining, nextSet)...)
	}
	return result
}

func permute(nums []int) [][]int {
	return compute(nums, []int{})
}
