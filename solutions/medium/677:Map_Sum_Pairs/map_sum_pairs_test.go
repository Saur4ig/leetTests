package _77_Map_Sum_Pairs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MapSum
	}{
		{
			name: "Example 1",
			want: MapSum{values: make(map[string]int)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Constructor()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name       string
		obj        MapSum
		inputKey   string
		inputValue int
		want       MapSum
	}{
		{
			name:       "Example 1",
			obj:        MapSum{values: map[string]int{"teserakt": 2}},
			inputKey:   "test1",
			inputValue: 3,
			want:       MapSum{values: map[string]int{"teserakt": 2, "test1": 3}},
		},
		{
			name:       "Example 2",
			obj:        MapSum{values: map[string]int{"sonet": 4}},
			inputKey:   "sonet",
			inputValue: 1,
			want:       MapSum{values: map[string]int{"sonet": 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.obj.Insert(tt.inputKey, tt.inputValue)
			require.Equal(t, tt.want, tt.obj)
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		obj    MapSum
		prefix string
		want   int
	}{
		{
			name:   "Example 1",
			obj:    MapSum{values: map[string]int{"apple": 3, "app": 2}},
			prefix: "ap",
			want:   5,
		},
		{
			name:   "Example 2",
			obj:    MapSum{values: map[string]int{"apple": 3, "app": 2, "router": 11, "postapp": 4}},
			prefix: "post",
			want:   4,
		},
		{
			name:   "Example 3",
			obj:    MapSum{values: map[string]int{"test": 1, "sss": 2, "vasdf": 3}},
			prefix: "",
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.obj.Sum(tt.prefix)
			require.Equal(t, tt.want, got)
		})
	}
}
