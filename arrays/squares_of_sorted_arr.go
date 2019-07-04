package arrays

import (
	"math"
	"math/rand"
)

// 977. Squares of a Sorted Array
// Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.
// todo: bad solution, rewrite with no sorting
// link - https://leetcode.com/problems/squares-of-a-sorted-array/
func sortedSquares(A []int) []int {
	tempArr := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		tempArr[i] = int(math.Pow(float64(A[i]), 2))
	}
	return quickSort(tempArr)
}

func quickSort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	l, r := 0, len(A)-1
	pivot := rand.Int() % len(A)
	A[pivot], A[r] = A[r], A[pivot]
	for i := range A {
		if A[i] < A[r] {
			A[l], A[i] = A[i], A[l]
			l++
		}
	}
	A[l], A[r] = A[r], A[l]

	quickSort(A[:l])
	quickSort(A[l+1:])
	return A
}
