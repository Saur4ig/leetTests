package _18

import (
	"testing"
)

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			name:  "Example 1",
			words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
			want:  16,
		},
		{
			name:  "Example 2",
			words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
			want:  4,
		},
		{
			name:  "Example 3",
			words: []string{"a", "aa", "aaa", "aaaa"},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
