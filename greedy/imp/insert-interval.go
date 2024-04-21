package greedy

/*

57. Insert Interval

You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Note that you don't need to modify intervals in-place. You can make a new array and return it.



Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

*/

func Insert(intervals [][]int, newInterval []int) [][]int {
	var left, right [][]int
	s := newInterval[0]
	e := newInterval[1]

	for _, v := range intervals {
		// check for intervals that less than newInterval
		if v[1] < s {
			left = append(left, v)
		} else if e < v[0] { // check for intervals that greater than newInterval
			right = append(right, v)
		} else { // the interval overlaps with new interval,so merge the inteval
			s = min(s, v[0])
			e = max(e, v[1])
		}
	}

	// fmt.Println(left,right, s, e)

	left = append(left, []int{s, e})

	return append(left, right...)
}
