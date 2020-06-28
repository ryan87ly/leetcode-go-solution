package lc437

import "testing"

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	level5 := &TreeNode{5, nil, nil}
	level4 := &TreeNode{4, nil, level5}
	level3 := &TreeNode{3, nil, level4}
	level2 := &TreeNode{2, nil, level3}
	level1 := &TreeNode{1, nil, level2}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{level1, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
