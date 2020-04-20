package lc85

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{
				[][]byte{
					[]byte{'1', '0', '1', '0', '0'},
					[]byte{'1', '0', '1', '1', '1'},
					[]byte{'1', '1', '1', '1', '1'},
					[]byte{'1', '0', '0', '1', '0'},
				},
			},
			6,
		},
		{
			"test2",
			args{
				[][]byte{
					[]byte{'1'},
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
