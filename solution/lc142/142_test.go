package lc142

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	node1 := ListNode{
		3,
		nil,
	}
	node2 := ListNode{
		2,
		nil,
	}
	node3 := ListNode{
		0,
		nil,
	}
	node4 := ListNode{
		-4,
		nil,
	}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node2

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"test1", args{&node1}, &node2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
