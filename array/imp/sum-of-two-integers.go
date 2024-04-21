package array

/*

371. Sum of Two Integers

Given two integers a and b, return the sum of the two integers without using the operators + and -.



Example 1:

Input: a = 1, b = 2
Output: 3
Example 2:

Input: a = 2, b = 3
Output: 5

*/

func getSum(a int, b int) int {
	keep := (a & b) << 1
	ans := a ^ b

	if keep == 0 {
		return ans
	}

	return getSum(ans, keep)
}
