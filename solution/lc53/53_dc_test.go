package lc53

import "testing"

func Test_maxSubArray2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}
