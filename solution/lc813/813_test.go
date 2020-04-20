package lc813

import "testing"

func Test_largestSumOfAverages(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{[]int{9, 1, 2, 3, 9}, 3}, 20.0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumOfAverages(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("largestSumOfAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
