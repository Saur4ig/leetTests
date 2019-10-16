package _77_Map_Sum_Pairs

import (
	"strings"
)

type MapSum struct {
	values map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{values: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.values[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	if prefix == "" || prefix == " " {
		return 0
	}
	var sum int
	for key, val := range this.values {
		if strings.HasPrefix(key, prefix) {
			sum += val
		}
	}
	return sum
}
