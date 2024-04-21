package dynamicProg

import "fmt"

/*
377. Combination Sum IV

Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.

The test cases are generated so that the answer can fit in a 32-bit integer.



Example 1:

Input: nums = [1,2,3], target = 4
Output: 7
Explanation:
The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Note that different sequences are counted as different combinations.
Example 2:

Input: nums = [9], target = 3
Output: 0


Constraints:

1 <= nums.length <= 200
1 <= nums[i] <= 1000
All the elements of nums are unique.
1 <= target <= 1000


Follow up: What if negative numbers are allowed in the given array?
 How does it change the problem? What limitation we need to add to the question to allow negative numbers?
*/

// no need keep track of nums, if we keep track of nums then permutation of those number will
// not be considered
// eg 1,1,2 will be consider but not 2,1,1 or 1,1,2
// exactly like coin change2 problem
func combinationSumUnique(nums []int, target int) int {
	dp := make([][]int, 0, target+1)
	// initialise the dp array
	for row := 0; row <= target; row++ {
		tmp := make([]int, 0, len(nums)+1)
		for col := 0; col <= len(nums); col++ {
			if row == 0 { // row ==0 and col ==0 will be handled here
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, 0)
			}
		}
		dp = append(dp, tmp)
	}
	fmt.Println(dp)

	for row := 1; row <= target; row++ {
		for col := 1; col <= len(nums); col++ {
			if row-nums[col-1] >= 0 {
				dp[row][col] += dp[row][col-1] + dp[row-nums[col-1]][col]
			} else {
				dp[row][col] += dp[row][col-1]
			}
		}

	}

	return dp[target][len(nums)]
}

func combinationSum4(nums []int, target int) int {
	return combinationSum(nums, target)
}

func combinationSum(nums []int, target int) int {
	dp := make([]int, target+1, target+1)
	dp[0] = 1

	for row := 1; row <= target; row++ {
		for col := 0; col < len(nums); col++ {
			if row-nums[col] >= 0 {
				dp[row] += dp[row-nums[col]]
			}
		}
	}

	return dp[target]
}

// all possible combination
// no need keep track of nums in dp array(dp[target][nums]), if we keep track of nums
// then permutation of those number will not be considered
// eg 1,1,2 will be consider but not 2,1,1 or 1,1,2
//sum(target) = sum(target - nums[i]) for all i -> 1 to nums
