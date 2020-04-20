package lc44

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{"aa", "a"}, false},
		{"test2", args{"aa", "*"}, true},
		{"test3", args{"cb", "?a"}, false},
		{"test4", args{"adceb", "*a*b"}, true},
		{"test5", args{"acdcb", "a*c?b"}, false},
		{"test6", args{"aa", "aa*"}, true},
		{"test7", args{"", "**"}, true},
		{"test8", args{"abefcdgiescdfimde", "ab*cd?i*de"}, true},
		{"test9", args{"aaaa", "***a"}, true},
		{"test10", args{"mississippi", "m*issip*"}, true},
		{"test11", args{"abcde", "*?*?*?*?"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
