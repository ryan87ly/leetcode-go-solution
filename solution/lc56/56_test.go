package lc56

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{
				[][]int{
					[]int{1, 3},
					[]int{2, 6},
					[]int{8, 10},
					[]int{15, 18},
				},
			},
			[][]int{
				[]int{1, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
