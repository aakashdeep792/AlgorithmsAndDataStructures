package graph

/*

128. Longest Consecutive Sequence

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.



Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

*/

func updateLenUtil(m map[int]int, val int) int {
	if _, ok := m[val]; !ok {
		return 0
	}

	if m[val] > 0 {
		return m[val]
	}

	m[val] = 1 + updateLenUtil(m, val+1)
	return m[val]
}

func updateLength(nums []int) int {
	has := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		has[nums[i]] = 0
	}

	for i := 0; i < len(nums); i++ {
		updateLenUtil(has, nums[i])
	}
	ans := 0
	for _, v := range has {
		ans = max(ans, v)
	}

	return ans
}

func LongestConsecutive(nums []int) int {
	return updateLength(nums)
}
