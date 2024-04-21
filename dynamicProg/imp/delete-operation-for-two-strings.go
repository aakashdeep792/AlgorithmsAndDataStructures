package dynamicProg

import "math"

/*

583. Delete Operation for Two Strings

Given two strings word1 and word2, return the minimum number of steps required to make word1 and word2 the same.

In one step, you can delete exactly one character in either string.



Example 1:

Input: word1 = "sea", word2 = "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
Example 2:

Input: word1 = "leetcode", word2 = "etco"
Output: 4


*/

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	dp := make([][]int, 0, l1+1)

	for i := 0; i <= l1; i++ {
		tmp := make([]int, 0, l2+1)
		for j := 0; j <= l2; j++ {
			if i == 0 || j == 0 {
				// if one word == "", all other characters of other word need to be deleted
				tmp = append(tmp, i+j)
			} else {
				tmp = append(tmp, math.MaxInt)
			}
		}
		dp = append(dp, tmp)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// similar to edit distance
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[l1][l2]

}

// add a space optimized solution of it
