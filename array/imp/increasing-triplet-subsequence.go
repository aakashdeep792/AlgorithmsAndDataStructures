package array

import "math"

func IncreasingTriplet(nums []int) bool {
	a, b := math.MaxInt, math.MaxInt
	for _, v := range nums {
		if v < a {
			a = v
		} else if v < b && v != a {
			b = v
		} else if v > b {
			return true
		}
	}
	return false
}
