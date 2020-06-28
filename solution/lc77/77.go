package lc77

func copyList(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

func backTrack(cur []int, used []bool, startIdx int, k int) [][]int {
	if len(cur) == k {
		return [][]int{cur}
	}
	result := make([][]int, 0)
	for i := startIdx; i < len(used); i++ {
		if !used[i] {
			next := append(copyList(cur), i+1)
			used[i] = true
			result = append(result, backTrack(next, used, i+1, k)...)
			used[i] = false
		}
	}
	return result
}

func combine(n int, k int) [][]int {
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		used[i] = false
	}
	return backTrack([]int{}, used, 0, k)
}
