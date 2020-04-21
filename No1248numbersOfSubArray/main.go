package No1248numbersOfSubArray

// 假如有 n 个奇数，那么当我们找的两端为奇数的优美数组，首位是第 i 个奇数的时候，末位必然是第 i+k-1 个奇数。

// 而第 i+k-1 个奇数，尾部可接的可能性数量，就是第 i+k 个奇数的前缀数量

func numberOfSubarrays(nums []int, k int) int {
	odd := make([]int, len(nums)+1)
	oddNum, prefixLen := 0, 0
	ans := 0
	for i := 0; i <= len(nums); i++ {
		prefixLen += 1 //自己也要算进去
		if i == len(nums) || nums[i]%2 == 1 {
			odd[oddNum] = prefixLen
			if oddNum >= k {
				ans += odd[oddNum-k] * odd[oddNum]
			}
			oddNum += 1
			prefixLen = 0
		}
	}
	return ans
}

