package strings

/*

3. Longest Substring Without Repeating Characters

Given a string s, find the length of the longest
substring
 without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	ans := 0
	freq := make(map[byte]int)
	for si, e := 0, 0; e < l; e++ {
		freq[s[e]]++

		for ; freq[s[e]] > 1; si++ {
			freq[s[si]]--

		}
		ans = max(ans, e-si+1)
	}

	return ans
}
