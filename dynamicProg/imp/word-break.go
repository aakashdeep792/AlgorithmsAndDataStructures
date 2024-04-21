package dynamicProg

import "fmt"

/*

139. Word Break

Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.



Example 1:

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false

*/

// solution 1
func wordBreak1(s string, wordDict []string) bool {
	dict := make(map[string]struct{})
	l := len(s)
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = struct{}{}
	}
	var dp [][]bool
	for i := 0; i < l; i++ {
		tmp := make([]bool, l, l)
		dp = append(dp, tmp)
		if _, ok := dict[string(s[i])]; ok {
			dp[i][i] = true
		}

	}

	for sl := 2; sl <= len(s); sl++ {
		for l, m := 0, sl-1; m < len(s); l, m = l+1, m+1 {
			// check if substring exist in dict if yes update dp[l][m]=true
			// if not then break the substring and check
			tmp := s[l : m+1]
			// fmt.Println(sl,tmp)
			if _, ok := dict[tmp]; ok {
				dp[l][m] = true
			}

			for k := l; k < m; k++ {
				dp[l][m] = dp[l][m] || (dp[l][k] && dp[k+1][m])
			}
		}
	}
	// fmt.Println(dp)
	return dp[0][l-1]
}

// solution2
func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	dp := make([]bool, l+1, l+1)
	dp[0] = true
	dict := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}

	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			dp[i+1] = dp[i+1] || (dp[j] && dict[s[j:i+1]])
		}
	}
	fmt.Println(dp)

	return dp[l]
}
