package lc56

import "sort"

type sortBy [][]int

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	if len(intervals) == 1 {
		return intervals
	}
	arr := sortBy(intervals)
	sort.Sort(arr)
	result := make([][]int, 0)
	curStart, curEnd := arr[0][0], arr[0][1]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if curEnd < interval[0] {
			result = append(result, []int{curStart, curEnd})
			curStart, curEnd = interval[0], interval[1]
		} else if curEnd <= interval[1] {
			curEnd = interval[1]
		}
	}
	result = append(result, []int{curStart, curEnd})
	return result
}
