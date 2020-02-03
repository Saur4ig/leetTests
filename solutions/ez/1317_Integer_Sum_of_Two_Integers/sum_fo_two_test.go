package _317_Integer_Sum_of_Two_Integers

import (
	"reflect"
	"testing"
)

func Test_getNoZeroIntegers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "Example 1",
			n:    2,
			want: []int{1, 1},
		},
		{
			name: "Example 2",
			n:    11,
			want: []int{2, 9},
		},
		{
			name: "Example 3",
			n:    10000,
			want: []int{1, 9999},
		},
		{
			name: "Example 4",
			n:    69,
			want: []int{1, 68},
		},
		{
			name: "Example 5",
			n:    1010,
			want: []int{11, 999},
		},
		{
			name: "Example 6",
			n:    20302,
			want: []int{2121, 18181},
		},
		{
			name: "Example 7",
			n:    10099,
			want: []int{111, 9988},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNoZeroIntegers(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNoZeroIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroIn(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{
			name:   "Example 1",
			number: 1234,
			want:   false,
		},
		{
			name:   "Example 2",
			number: 1,
			want:   false,
		},
		{
			name:   "Example 3",
			number: 10,
			want:   true,
		},
		{
			name:   "Example 4",
			number: 100,
			want:   true,
		},
		{
			name:   "Example 5",
			number: 109532,
			want:   true,
		},
		{
			name:   "Example 6",
			number: 350335,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroIn(tt.number); got != tt.want {
				t.Errorf("isZeroIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
