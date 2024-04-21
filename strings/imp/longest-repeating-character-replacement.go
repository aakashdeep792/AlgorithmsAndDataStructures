package strings

/*

424. Longest Repeating Character Replacement

You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.



Example 1:

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
Example 2:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.

*/

func maxFreq(freq []int) int {
	ans := 0
	for _, v := range freq {
		ans = max(ans, v)
	}
	return ans
}

func characterReplacement(s string, k int) int {
	freq := make([]int, 26)
	count := 0
	ans := 0
	for i, j := 0, 0; i < len(s); i++ {
		freq[s[i]-'A']++
		count++
		// remove the char one by one ,check if it satishfy the condition,
		// if yes than store the count
		for count-maxFreq(freq) > k {
			freq[s[j]-'A']--
			count--
			j++
		}
		ans = max(ans, count)
	}
	return ans
}
