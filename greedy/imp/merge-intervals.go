package greedy

import (
	"fmt"
	"sort"
)

/*
56. Merge Intervals

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.



Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

*/

func mergeUtil(intervals [][]int) [][]int {
	// i<j so interval[i][1]<interval[j][1] for sorted
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	fmt.Println(intervals)
	var ans [][]int
	for _, v := range intervals {
		if len(ans) == 0 {
			ans = append(ans, v)
		} else if ans[len(ans)-1][1] < v[0] { // non overlap case
			ans = append(ans, v)
		} else { //overlap case
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], v[1])

		}

	}

	return ans
}

func merge(intervals [][]int) [][]int {
	return mergeUtil(intervals)
}
