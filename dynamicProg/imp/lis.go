package dynamicProg

/*

300. Longest Increasing Subsequence

Given an integer array nums, return the length of the longest strictly increasing
subsequence
.


// lower bound return the index of  next smallest number
// just greater than or equal to target number.
func lowerBound(nums []int, x int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if x <= nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func lisGreedy(nums []int) int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > ans[len(ans)-1] {
			ans = append(ans, nums[i])
		} else {
			l := lowerBound(ans, nums[i])
			ans[l] = nums[i]
		}
	}

	return len(ans)
}

func lengthOfLIS(nums []int) int {
	return lisGreedy(nums)
}

Input: nums = [7,7,7,7,7,7,7]
Output: 1

*/

// lower bound return the index of  next smallest number
// just greater than or equal to target number.
func lowerBound(nums []int, x int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if x <= nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func lisGreedy(nums []int) int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > ans[len(ans)-1] {
			ans = append(ans, nums[i])
		} else {
			l := lowerBound(ans, nums[i])
			ans[l] = nums[i]
		}
	}

	return len(ans)
}

func lengthOfLIS(nums []int) int {
	return lisGreedy(nums)
}
