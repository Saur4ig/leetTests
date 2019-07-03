package arrays

// 832. Flipping an Image
// Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.
// link - https://leetcode.com/problems/flipping-an-image/
func FlipAndInvertImage(A [][]int) [][]int {
	length := len(A)
	res := make([][]int, 0)
	for _, innerArray := range A {
		newInner := make([]int, 0)
		for i := length - 1; i >= 0; i-- {
			if innerArray[i] == 0 {
				newInner = append(newInner, 1)
			} else {
				newInner = append(newInner, 0)
			}
		}
		res = append(res, newInner)
	}
	return res
}
