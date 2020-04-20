package lc312

import "testing"

func Test_maxCoins(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {"test1", args{[]int{3, 1, 5, 8}}, 167},
		{"test2", args{[]int{5, 8}}, 48},
		// {"test3", args{[]int{5, 2, 3}}, 50},
		{"test4", args{[]int{5}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
