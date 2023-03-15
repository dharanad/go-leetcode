package main

import "sort"

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	// Find Level Sum
	// Check valid
	// Sort
	// return l - k value
	var q []*TreeNode
	q = append(q, root)
	var levelSumArr []int64
	for len(q) > 0 {
		var levelSum int64
		levelCount := len(q)
		for i := 0; i < levelCount; i++ {
			curr := q[0]
			q = q[1:]
			levelSum += int64(curr.Val)
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		levelSumArr = append(levelSumArr, levelSum)
	}
	if len(levelSumArr) < k {
		return -1
	}
	sort.Slice(levelSumArr, func(i, j int) bool {
		return levelSumArr[i] > levelSumArr[j]
	})
	return levelSumArr[k-1]
}
