package strings

import (
	"sort"
	"strings"
)

/*

49. Group Anagrams

Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]

*/

func sortString(s string) string {
	tmp := strings.Split(s, "")
	sort.Strings(tmp)
	return strings.Join(tmp, "")
}

// T=O(NKlogK)
func groupAnagrams(strs []string) [][]string {
	g := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		tmp := sortString(strs[i])

		g[tmp] = append(g[tmp], strs[i])

	}

	ans := make([][]string, 0, len(strs))
	for _, v := range g {
		ans = append(ans, v)
	}

	return ans
}

// sol2 T= O(NK) using hash map
