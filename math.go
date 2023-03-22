package main

import (
	"sort"
)

func CountDigit(x int) int {
	digitCount := 0
	for x > 0 {
		digitCount++
		x /= 10
	}
	return digitCount
}

func findFarmland(land [][]int) [][]int {
	rows := len(land)
	if rows == 0 {
		return nil
	}
	cols := len(land[0])
	var ans [][]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// If forest land or visited then skip
			if land[i][j] == 0 {
				continue
			}

			nj := j
			// Right most boundary
			for nj < cols && land[i][nj] == 1 {
				nj++
			}

			ni := i
			for ni < rows && land[ni][j] == 1 {
				ni++
			}

			// Handle zero index case
			if nj != 0 {
				nj -= 1
			}

			if ni != 0 {
				ni -= 1
			}
			auxAns := []int{i, j, ni, nj}
			ans = append(ans, auxAns)
			// Mark the rectangle visited
			for k := i; k <= ni; k++ {
				for l := j; l <= nj; l++ {
					land[k][l] = 0
				}
			}
		}
	}
	return ans
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	// Since the array are already sorted, we can use this property
	// We will extend the merge algorithm of the merge sort to solve this problem
	var res [][]int
	n, m := len(nums1), len(nums2)
	i, j := 0, 0
	for i < n && j < m {
		id1 := nums1[i][0]
		id2 := nums2[j][0]
		if id1 < id2 {
			res = append(res, nums1[i])
			i++
		} else if id1 > id2 {
			res = append(res, nums2[j])
			j++
		} else {
			nums2[j][1] += nums1[i][1]
			res = append(res, nums2[j])
			i++
			j++
		}
	}
	if i < n {
		res = append(res, nums1[i:]...)
	}
	if j < m {
		res = append(res, nums2[j:]...)
	}
	return res
}

func divisibilityArray(word string, m int) []int {
	arrLen := len(word)
	ans := make([]int, arrLen)
	var r int64 = 0
	for idx := range word {
		r = (r*10 + int64(word[idx]-'0')) % int64(m)
		if r == 0 {
			ans[idx] = 1
		}
	}
	return ans
}

func maxNumOfMarkedIndices(nums []int) int {
	arrLen := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	i := 0
	for idx := arrLen - arrLen/2; idx < arrLen; idx++ {
		if 2*nums[i] <= nums[idx] {
			i++
		}
	}
	return 2 * i
}

func twoSum(nums []int, target int) []int {
	valIdxMap := make(map[int]int)
	for idx, val := range nums {
		if otherIdx, ok := valIdxMap[target-val]; ok {
			return []int{otherIdx, idx}
		}
		valIdxMap[val] = idx
	}
	return []int{}
}
