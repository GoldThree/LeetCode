package No4trap

func Trap3(height []int) int {  // 暴力法
	var ans int

	cols := len(height)

	for i := 1; i < cols-1; i++{
		var maxLeft int
		var maxRight int
		var lower int

		for j := 0; j < i; j++{
			maxLeft = Max(maxLeft, height[j])
		}

		for j := i + 1; j < cols; j++{
			maxRight = Max(maxRight, height[j])
		}

		lower = Min(maxLeft, maxRight)

		if lower > height[i]{
			ans = ans + lower - height[i]
		}
	}

	return ans
}

func Trap2(height []int) int {  // 暴力法基础上 动态规划
	var ans int

	cols := len(height)
	maxLeft := make([]int, cols)
	maxRight := make([]int, cols)

	for i := 1; i < cols; i++{
		maxLeft[i] = Max(maxLeft[i-1], height[i-1])
	}

	for i := cols - 2; i >= 0; i--{  // 特别注意的是 maxRight[i] 依赖 maxRight[i+1] 所以i 从大到小遍历
		maxRight[i] = Max(maxRight[i+1], height[i+1])
	}

	for i := 1; i < cols-1; i++{
		var lower int

		lower = Min(maxLeft[i], maxRight[i])

		if lower > height[i]{
			ans = ans + lower - height[i]
		}
	}

	return ans
}

func trap(height []int) int {  // 双指针
	var ans int
	cols := len(height)

	var left  = 1
	var right = cols - 2

	var maxLeft int
	var maxRight int

	for left <= right{
		maxLeft = Max(maxLeft, height[left - 1])
		maxRight = Max(maxRight, height[right + 1])
		if maxLeft < maxRight{
			var lower = maxLeft
			if lower > height[left]{
				ans = ans + lower - height[left]
			}
			left++
		}else{
			var lower = maxRight
			if lower > height[right]{
				ans = ans + lower - height[right]
			}
			right--
		}
	}

	return ans
}

func trap4(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				//设置左边最高柱子
				leftMax = height[left]
			} else {
				//右边必定有柱子挡水，所以，遇到所有值小于等于leftMax的，全部加入水池
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				//设置右边最高柱子
				rightMax = height[right]
			} else {
				//左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}

func Max(a,b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min(a,b int) int {
	if a >= b {
		return b
	}
	return a
}
