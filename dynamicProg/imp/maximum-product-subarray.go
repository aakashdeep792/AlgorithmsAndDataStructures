package dynamicProg

/*
152. Maximum Product Subarray

Given an integer array nums, find a
subarray
 that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.



Example 1:

Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

*/

func maxThree(a, b, c int) int {
	return max(a, max(b, c))
}

func minThree(a, b, c int) int {
	return min(a, min(b, c))
}

// using kadance
func MaxProduct(nums []int) int {
	mini := nums[0]
	maxi := nums[0]
	ans := nums[0]
	for r := 1; r < len(nums); r++ {
		//
		tmp := maxThree(nums[r], maxi*nums[r], mini*nums[r])
		mini = minThree(nums[r], maxi*nums[r], mini*nums[r])
		maxi = tmp
		ans = max(ans, maxi)
	}

	return ans
}
