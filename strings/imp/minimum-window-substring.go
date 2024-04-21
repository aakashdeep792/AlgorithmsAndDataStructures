package strings

import "math"

/*

76. Minimum Window Substring

Given two strings s and t of lengths m and n respectively, return the minimum window
substring
 of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.


*/

func match(str, pat []int) bool {
	for i, v := range pat {
		if str[i] < v {
			return false
		}
	}
	return true
}

func getMinWindow(s, t string) string {
	ll, ans := -1, math.MaxInt
	freqS, freqP := make([]int, 256), make([]int, 256)
	for i := 0; i < len(t); i++ {
		freqP[t[i]]++
	}

	for l, r := 0, 0; r < len(s); r++ {
		freqS[s[r]]++
		for match(freqS, freqP) && l <= r {
			if ans > r-l+1 {
				ll, ans = l, r-l+1
			}

			freqS[s[l]]--
			l++
		}

	}

	if ans == math.MaxInt {
		return ""
	}
	return s[ll : ll+ans]
}

func minWindow(s string, t string) string {
	return getMinWindow(s, t)
}
