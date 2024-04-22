package array

import "fmt"

/*

238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

*/

func productExceptSelfUtil(nums []int) []int {
	var (
		prod      int64 = 1
		ans       []int
		countZero int
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			countZero++
			continue
		}

		prod *= int64(nums[i])
	}

	prev := 1
	for _, val := range nums {
		// when array contain zeros
		if countZero > 0 {
			tmp := 0
			if countZero == 1 && val == 0 {
				tmp = int(prod)
			}
			ans = append(ans, tmp)
			continue
		}

		prod = prod * int64(prev) / int64(val)
		fmt.Println(prod)
		ans = append(ans, int(prod))
		prev = val
	}
	return ans
}

// Using prefix and suffix product
func productExceptSelfUtil2(nums []int) []int {
	ans := make([]int, 0, len(nums))
	prod := 1
	// calculate the prefix sum  ans[i]= num[0]*...*num[i-1]
	for _, val := range nums {
		ans = append(ans, prod)
		prod *= val
	}
	prod = 1
	// this creat product of prefix and suffix .
	// ans[i]= num[0]*...*num[i-1]*num[i+1]*...*num[len-1]
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= prod
		prod *= nums[i]
	}
	return ans
}
func ProductExceptSelf(nums []int) []int {
	// return productExceptSelfUtil(nums)
	return productExceptSelfUtil2(nums)
}
