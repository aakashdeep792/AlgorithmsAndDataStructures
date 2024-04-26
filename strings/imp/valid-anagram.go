package strings

/*

242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false

*/

func isMatch(freqS, freqT []int) bool {
	for i := 0; i < len(freqS); i++ {
		if freqS[i] != freqT[i] {
			return false
		}
	}

	return true
}
func IsAnagram(s string, t string) bool {
	freqS := make([]int, 26)
	freqP := make([]int, 26)

	for i := 0; i < len(s); i++ {
		freqS[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		freqP[t[i]-'a']++
	}

	return isMatch(freqS, freqP)
}
