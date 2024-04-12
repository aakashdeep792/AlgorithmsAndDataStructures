package dynamicProg

/*
1143. Longest Common Subsequence

Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

 Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

*/

func lcs(str1 string, str2 string) int {
	infinite := 10000000000
	l1 := len(str1)
	l2 := len(str2)
	dp := make([][]int, 0, l1+1)
	// populate dp array
	for row := 0; row <= l1; row++ {
		tmp := make([]int, 0, l2+1)
		for col := 0; col <= l2; col++ {
			if row == 0 || col == 0 {
				tmp = append(tmp, 0)
			} else {
				tmp = append(tmp, infinite)
			}
		}
		dp = append(dp, tmp)
	}

	for row := 1; row <= l1; row++ {
		for col := 1; col <= l2; col++ {
			if str1[row-1] == str2[col-1] {
				dp[row][col] = 1 + dp[row-1][col-1]
			} else {
				dp[row][col] = max(dp[row-1][col], dp[row][col-1])
			}
		}
	}
	return dp[l1][l2]

}

func longestCommonSubsequence(text1 string, text2 string) int {
	return lcs(text1, text2)
}
