package binarySearch

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if target <= nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if l >= 0 && l < len(nums) && nums[l] == target {
		return l
	} else {
		return -1
	}

}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if r >= 0 && r < len(nums) && nums[r] == target {
		return r
	} else {
		return -1
	}
}

func SearchRange(nums []int, target int) []int {
	l := lowerBound(nums, target)
	r := upperBound(nums, target)

	return []int{l, r}
}
