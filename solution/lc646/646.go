package lc646

import "sort"

type pairs [][]int

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findLongestChain(inputPairs [][]int) int {
	ps := pairs(inputPairs)
	sort.Sort(ps)
	res, right := 1, ps[0][1]
	for i := 1; i < len(ps); i++ {
		if ps[i][0] > right {
			res, right = res+1, ps[i][1]
		}
	}
	return res
}
