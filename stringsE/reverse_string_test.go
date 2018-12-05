package stringsE

import "testing"

func TestReverseString(t *testing.T) {
	input := "hello"
	output := "olleh"
	result := ReverseString(input)

	if result != output {
		t.Errorf("incorrect, got: %s, want: %s.", result, output)
	}
}
