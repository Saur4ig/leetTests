package _309_Decrypt_String

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_freqAlphabets(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "10#11#12",
			want: "jkab",
		},
		{
			name: "Example 2",
			s:    "1326#",
			want: "acz",
		},
		{
			name: "Example 3",
			s:    "25#",
			want: "y",
		},
		{
			name: "Example 4",
			s:    "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
			want: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name: "Example 5",
			s:    "1111111",
			want: "aaaaaaa",
		},
		{
			name: "Example 6",
			s:    "26#1",
			want: "za",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := freqAlphabets(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}
