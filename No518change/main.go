package No518change

// 动态规划 重点看

func change(amount int, coins []int) int {

	dp := make([]int, amount + 1)

	for key, _ := range dp {
		dp[key] = 0
	}
	dp[0] = 1
	for _, coin := range coins {
		for i:= coin;i<amount+1;i++ {
			// fmt.Println(i, coin, dp)
			dp[i] += dp[i-coin]
		}
	}
	// fmt.Println(dp)

	return dp[amount]
}
