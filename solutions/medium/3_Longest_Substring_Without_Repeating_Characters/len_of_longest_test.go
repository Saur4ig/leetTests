package __Longest_Substring_Without_Repeating_Characters

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "Example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "Example 3",
			s:    "",
			want: 0,
		},
		{
			name: "Example 4",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "Example 5",
			s:    "au",
			want: 2,
		},
		{
			name: "Example 6",
			s:    "c",
			want: 1,
		},
		{
			name: "Example 7",
			s:    "dvdf",
			want: 3,
		},
		{
			name: "Example 8",
			s:    "abba",
			want: 2,
		},
		{
			name: "Example 9",
			s:    "uqinntq",
			want: 4,
		},
		{
			name: "Example 10",
			s:    "aabaab!bb",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
