package _87_First_Unique_Character_in_a_String

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example 1",
			input: "leetcode",
			want:  0,
		},
		{
			name:  "Example 2",
			input: "loveleetcode",
			want:  2,
		},
		{
			name:  "Example 3",
			input: "aaaaaaa",
			want:  -1,
		},
		{
			name:  "Example 4",
			input: "llooppssttk",
			want:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstUniqChar(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
