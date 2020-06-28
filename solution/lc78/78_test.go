package lc78

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{[]int{9, 0, 3, 5, 7}},
			[][]int{
				[]int{3},
				[]int{1},
				[]int{2},
				[]int{1, 2, 3},
				[]int{1, 3},
				[]int{2, 3},
				[]int{1, 2},
				[]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
