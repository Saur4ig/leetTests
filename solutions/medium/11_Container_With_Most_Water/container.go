package _1_Container_With_Most_Water

// timeout(too slow)
func maxArea(height []int) int {
	max := 0
	for i, v := range height {
		for j := len(height) - 1; j > i; j-- {
			curr := (j - i) * getMin(v, height[j])
			if curr > max {
				max = curr
			}
		}
	}
	return max
}

// Runtime: 148 ms, faster than 21.06% of Go online submissions for Container With Most Water.
// Memory Usage: 9 MB, less than 44.66% of Go online submissions for Container With Most Water.
func maxArea2(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		curr := (right - left) * getMin(height[left], height[right])
		if curr > max {
			max = curr
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
