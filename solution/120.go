package main

import (
	"math"
)

func min(nums ...int) int {
	curMin := math.MaxInt32
	for _, num := range nums {
		if curMin > num {
			curMin = num
		}
	}
	return curMin
}

func minimumTotal(triangle [][]int) int {
	layer := len(triangle)
	if layer == 1 {
		return triangle[0][0]
	}
	prevDP := make([]int, layer)
	curDP := make([]int, layer)

	// initilization
	curDP = triangle[0]
	prevDP = curDP

	for i := 1; i < layer; i++ {
		curDP = make([]int, layer)
		for j := 0; j <= i; j++ {
			curElem := triangle[i][j]
			if j == 0 {
				curDP[j] = prevDP[0] + curElem
			} else if j == i {
				curDP[j] = prevDP[j-1] + curElem
			} else {
				curDP[j] = min(prevDP[j-1], prevDP[j]) + curElem
			}
		}
		prevDP = curDP
	}
	return min(curDP...)
}

// func main() {
// 	fmt.Println(minimumTotal(
// 		[][]int{
// 			[]int{2},
// 			[]int{3, 4},
// 			[]int{6, 5, 7},
// 			[]int{4, 1, 8, 3},
// 		},
// 	))
// }
