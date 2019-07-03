package arrays

// 961. N-Repeated Element in Size 2N Array
// In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.
// Return the element repeated N times.
// link - https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
func repeatedNTimes(A []int) int {
	nums := make(map[int]bool)
	for _, val := range A {
		if _, ok := nums[val]; ok {
			return val
		}
		nums[val] = true
	}
	return -1
}
