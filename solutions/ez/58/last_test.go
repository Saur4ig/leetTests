package _8

import (
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "Hello World",
			want: 5,
		},
		{
			name: "Example 2",
			s:    "   fly me   to   the moon  ",
			want: 4,
		},
		{
			name: "Example 3",
			s:    "luffy is still joyboy",
			want: 6,
		},
		{
			name: "Example 4",
			s:    "a",
			want: 1,
		},
		{
			name: "Example 5",
			s:    "    day",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
