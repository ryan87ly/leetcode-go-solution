package main

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{[]int{23, 2, 6, 4, 7}, 6}, want: true},
		{name: "test2", args: args{[]int{0, 0}, 0}, want: true},
		{name: "test3", args: args{[]int{0, 1, 0}, 0}, want: false},
		{name: "test4", args: args{[]int{5, 0, 0}, 0}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
