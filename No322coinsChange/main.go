package No322coinsChange

//假设最小组合中最后一枚硬币是c：
//
//c 可能是 coins 中任何一个；
//去除 c 后剩下的部分，一定也是当前总额的一个最小组合（否则加上c不可能构成最小组合）
//或者用以下思路：
//如果 dp[i] 表示组成金额i的最小组合，dp[i]+1 一定是组成金额 i+c 的最小组合。

// func coinChange(coins []int, amount int) int {

//     // // 数组大小为 amount + 1，初始值也为 amount + 1
//     // vector<int> dp(amount + 1, amount + 1);
//     // // base case
//     // dp[0] = 0;
//     // for (int i = 0; i < dp.size(); i++) {
//     //     // 内层 for 在求所有子问题 + 1 的最小值
//     //     for (int coin : coins) {
//     //         // 子问题无解，跳过
//     //         if (i - coin < 0) continue;
//     //         dp[i] = min(dp[i], 1 + dp[i - coin]);
//     //     }
//     // }
//     // return (dp[amount] == amount + 1) ? -1 : dp[amount];

//     dp := make([]int, amount+1)

//     for key, _ := range dp {
//         if key == 0 {
//             dp[key] = 0
//             continue
//         }
//         dp[key] = amount+1
//     }

//     length := amount+1

//     for i:= 0;i<length;i++ {
//         for _, coin := range coins {
//             if i-coin < 0 {
//                 continue
//             }
//             dp[i]= min(dp[i], 1+dp[i-coin])
//         }
//     }

//     if dp[amount] == amount+1 {
//         return -1
//     }

//     return dp[amount]
// }

// func min(a, b int) int {
//     if a>=b {
//         return b
//     }
//     return a
// }

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}

			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}

	return dp[amount]
}

