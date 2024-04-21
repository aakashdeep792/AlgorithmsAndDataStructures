package array

/*

217. Contains Duplicate

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

func ContainsDuplicate(nums []int) bool {
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
		if freq[nums[i]] > 1 {
			return true
		}
	}
	return false
}

/*
other solution could be
1. using two nested loop, and check for repeated element(t= O(N*N), S=O(1)).
2. sort the number and the check for repeated element(t= O(N*logN), S=O(1) - randamised quick sort)
3. Using set, compare the set length(sl) with number of element(i) processed
    if sl< i mean repeatation exit (t= O(N*logN), S=O(N)
4. sorted map (t= O(N*logN), S=O(N)
5 hash map (t= O(N), S=O(1)
*/
