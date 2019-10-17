package _108_Defanging_an_IP_Address

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefangIPaddr(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Example 1",
			input: "1.1.1.1",
			want:  "1[.]1[.]1[.]1",
		},
		{
			name:  "Example 2",
			input: "255.100.50.0",
			want:  "255[.]100[.]50[.]0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defangIPaddr(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
