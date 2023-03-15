package main

import "sort"

func maxScore(nums []int) int {
	// sort in desc order, so that larger number to beginning of the array
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	count := 0
	sum := 0
	for _, x := range nums {
		sum += x
		if sum > 0 {
			count++
		}
	}
	return count
}
