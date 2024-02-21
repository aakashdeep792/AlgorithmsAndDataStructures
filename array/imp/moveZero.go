/*

 Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


Follow up: Could you minimize the total number of operations done?

*/

package main

func moveZeroes(nums []int) {
	moveNumberToEnd(nums, 0)
}

func moveNumberToEnd(nums []int, val int) {
	if len(nums) == 0 {
		return
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}

		nums[count], nums[i] = nums[i], nums[count]
		count++
	}
}
