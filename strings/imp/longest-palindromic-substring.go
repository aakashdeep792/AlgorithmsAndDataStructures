package strings

/*
5. Longest Palindromic Substring

Given a string s, return the longest
palindromic

substring
 in s.



Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"

*/

func getPalingdromLen(s string, i, j int) (int, int) {
	// initially start with same index
	l := i
	r := i
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}

		l = i
		r = j
		i--
		j++
	}

	return l, r
}

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	ll, rr := 0, 0
	for i := 0; i < len(s)-1; i++ {
		//even len
		l, r := getPalingdromLen(s, i, i+1)
		if r-l > rr-ll {
			rr = r
			ll = l
		}
		// odd len
		l, r = getPalingdromLen(s, i, i)
		if r-l > rr-ll {
			rr = r
			ll = l
		}
	}
	return s[ll : rr+1]
}
