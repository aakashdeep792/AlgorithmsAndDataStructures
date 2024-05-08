package array

import "sort"

/*
15. 3Sum

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

*/

func threeSumUtil(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	var ans [][]int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i // store the last occurance of every number
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ans
		}
		for j := i + 1; j < len(nums); j++ {
			req := -1 * (nums[i] + nums[j])
			if k, ok := m[req]; ok && k > j {
				ans = append(ans, []int{nums[i], nums[j], req})
			}
			j = m[nums[j]]
		}
		i = m[nums[i]]
	}
	return ans
}

func ThreeSum(nums []int) [][]int {
	return threeSumUtil(nums)
}

// implement using binary search
