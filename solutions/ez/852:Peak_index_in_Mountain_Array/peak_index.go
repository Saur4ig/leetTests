package _52_Peak_index_in_Mountain_Array

// 852. Peak Index in a Mountain Array
// Let's call an array A a mountain if the following properties hold:
//
// A.length >= 3
// There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
// Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
// link - https://leetcode.com/problems/peak-index-in-a-mountain-array/
func peakIndexInMountainArray(A []int) int {
	maxPos := 0
	max := A[maxPos]
	for i, val := range A {
		if val > max {
			max = val
			maxPos = i
		}
	}
	return maxPos
}
