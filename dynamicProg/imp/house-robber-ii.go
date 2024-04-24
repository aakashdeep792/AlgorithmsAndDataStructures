package dynamicProg

/*
213. House Robber II

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.



Example 1:

Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
Example 2:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 3:

Input: nums = [1,2,3]
Output: 3

*/

// both logic are same
func rob2(nums []int) int {
	inc := make([]int, len(nums)+1)
	exc := make([]int, len(nums)+1)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 2 to n
	for i := 1; i <= len(nums); i++ {
		if i > 1 {
			inc[i] = exc[i-1] + nums[i-1]
		}

		exc[i] = max(inc[i-1], exc[i-1])
	}

	ans := max(inc[len(nums)], exc[len(nums)])
	// 1 to n-1
	for i := 1; i <= len(nums); i++ {
		if i < len(nums) {
			inc[i] = exc[i-1] + nums[i-1]
		}

		exc[i] = max(inc[i-1], exc[i-1])
	}
	ans = max(ans, exc[len(nums)])
	// if len(nums)==1{
	//     return max(inc[len(nums)],exc[len(nums)])
	// }

	// if last element is part of solution then subtract the nums[0] and return the value
	// otherwise exc[len(nums)] is the solution
	// if flag == true{
	//     return max(inc[len(nums)]-nums[0],exc[len(nums)])
	// }

	return ans
}

func robber(nums []int, s, e int) int {
	// no need to create a dp array as only previous value is need
	inc, exc := 0, 0
	for i := s; i <= e; i++ {
		inc, exc = exc+nums[i], max(exc, inc)
		// OR
		// tmp:= exc + nums[i]
		// exc = max(exc, inc)
		// inc = tmp

	}
	return max(inc, exc)
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// max (  robber(0 to n-2 ), robber(1 to n-1 )
	return max(robber(nums, 0, len(nums)-2), robber(nums, 1, len(nums)-1))
}
