package _89_Rotate_Array

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k = len(nums) - k%len(nums)
	copy(nums, append(nums[k:], nums[:k]...))
}
