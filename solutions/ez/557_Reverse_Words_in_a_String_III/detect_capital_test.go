package _57_Reverse_Words_in_a_String_III

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Example 1",
			input: "test test",
			want:  "tset tset",
		},
		{
			name:  "Example 2",
			input: "Let's take LeetCode contest",
			want:  "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
