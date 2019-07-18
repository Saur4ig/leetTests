package stringsE

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []string
	}{
		{
			name:  "Example 1",
			input: 15,
			want:  []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fizzBuzz(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
