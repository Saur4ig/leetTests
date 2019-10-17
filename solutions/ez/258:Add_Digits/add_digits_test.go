package _58_Add_Digits

import (
	"testing"
)

func TestAddDigits(t *testing.T) {
	input := 38
	output := 2
	result := AddDigits(input)

	if result != output {
		t.Errorf("incorrect, got: %d, want: %d.", result, output)
	}
}
