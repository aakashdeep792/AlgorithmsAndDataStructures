package dynamicProg

/*

322. Coin Change

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0

*/

func coinChange(coins []int, amount int) int {
	return coinChangeUtilSpaceOptimized(coins, amount)
}

func coinChangeUtil(coins []int, amount int) int {
	infinite := 100000000000
	dp := make([][]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, len(coins)+1)
		for j := 0; j <= len(coins); j++ {
			if i == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = infinite
			}
		}

	}

	for sum := 1; sum <= amount; sum++ {
		for c := 1; c <= len(coins); c++ {
			if sum-coins[c-1] >= 0 {
				dp[sum][c] = min(dp[sum][c], min(dp[sum-coins[c-1]][c]+1, dp[sum][c-1]))
			} else {
				dp[sum][c] = min(dp[sum][c], dp[sum][c-1])
			}
		}
	}
	if dp[amount][len(coins)] >= infinite {
		return -1
	}
	return dp[amount][len(coins)]

}

// can we optimize it for O(N) space
func coinChangeUtilSpaceOptimized(coins []int, amount int) int {
	infinite := 100000000000
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = infinite

	}

	for sum := 1; sum <= amount; sum++ {
		for c := 0; c < len(coins); c++ {
			if sum-coins[c] >= 0 {
				dp[sum] = min(dp[sum], dp[sum-coins[c]]+1)
			}
		}
	}

	if dp[amount] >= infinite {
		return -1
	}

	return dp[amount]
}
