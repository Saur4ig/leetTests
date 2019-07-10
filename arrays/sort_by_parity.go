package arrays

// 922. Sort Array By Parity II
// Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even.
// Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.
// link - https://leetcode.com/problems/sort-array-by-parity-ii/
func sortArrayByParityII(A []int) []int {
	nextE, nextO := 0, 1
	arr := make([]int, len(A))
	for _, val := range A {
		if val%2 == 0 {
			arr[nextE] = val
			nextE += 2
			continue
		}
		arr[nextO] = val
		nextO += 2
	}
	return arr
}
