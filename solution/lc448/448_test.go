package lc448

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
		{"test2", args{[]int{2, 1}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
