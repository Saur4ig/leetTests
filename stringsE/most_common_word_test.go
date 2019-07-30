package stringsE

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMostCommonWord(t *testing.T) {
	tests := []struct {
		name           string
		inputParagraph string
		inputBanned    []string
		want           string
	}{
		{
			name:           "Example 1",
			inputParagraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			inputBanned:    []string{"hit"},
			want:           "ball",
		},
		{
			name:           "Example 2",
			inputParagraph: "let 1, be`s, the'asd let 1 1 1 1",
			inputBanned:    []string{"let"},
			want:           "1",
		},
		{
			name:           "Example 3",
			inputParagraph: "Bob!",
			inputBanned:    []string{"let"},
			want:           "bob",
		},
		{
			name:           "Example 4",
			inputParagraph: "a, a, a, a, b,b,b,c, c",
			inputBanned:    []string{"a"},
			want:           "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mostCommonWord(tt.inputParagraph, tt.inputBanned)
			require.Equal(t, tt.want, got)
		})
	}
}
