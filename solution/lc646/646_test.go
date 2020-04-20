package lc646

import "testing"

func Test_findLongestChain(t *testing.T) {
	type args struct {
		inputPairs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"Test1",
			args{
				[][]int{
					{1, 2},
					{2, 3},
					{3, 4},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestChain(tt.args.inputPairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
