package code

import (
	"fmt"
	"sync"
)

func merge(nums []int, e1, e2 int) {
	a, b := 0, e1+1
	tmp := make([]int, 0, e1+1)
	for a <= e1 && b <= e2 {
		if nums[a] <= nums[b] {
			tmp = append(tmp, nums[a])
			a++
		} else {
			tmp = append(tmp, nums[b])
			b++
		}
	}

	if b > e2 {
		for ; a <= e1; a++ {
			tmp = append(tmp, nums[a])
		}
	}
	if a > e1 {
		for ; b <= e1; b++ {
			tmp = append(tmp, nums[b])
		}
	}

	for i := 0; i <= e2; i++ {
		nums[i] = tmp[i]
	}

}

func MergeSort(wg *sync.WaitGroup, nums []int) {
	defer wg.Done()
	// len is 1
	if len(nums) <= 1 {
		return
	}
	// len is 2
	if len(nums) == 2 {
		nums[0], nums[1] = min(nums[0], nums[1]), max(nums[0], nums[1])
		return
	}

	m := (len(nums) - 1) / 2

	var w2 sync.WaitGroup
	w2.Add(2)
	go MergeSort(&w2, nums[:m+1])
	go MergeSort(&w2, nums[m+1:])
	w2.Wait()

	merge(nums, m, len(nums)-1)
}

func SortDriver() string {

	nums := []int{5, 3, 2, 1, 4}

	var wg sync.WaitGroup
	wg.Add(1)
	go MergeSort(&wg, nums)
	wg.Wait()
	fmt.Println(nums)
	return "success"

}
