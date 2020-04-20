package lc546

import "testing"

func Test_removeBoxes(t *testing.T) {
	type args struct {
		boxes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.

		{"test1", args{[]int{1, 3, 2, 2, 2, 3, 4, 3, 1}}, 23},
		{"test2", args{[]int{1, 2}}, 2},
		{"test3", args{[]int{1, 2, 1}}, 5},
		{"test4", args{[]int{1, 1, 2, 1, 1}}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBoxes(tt.args.boxes); got != tt.want {
				t.Errorf("removeBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
