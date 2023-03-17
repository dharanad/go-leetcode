package main

func leftRigthDifference(nums []int) []int {
	arrLen := len(nums)
	ans := make([]int, arrLen)
	rightSum := 0
	for _, val := range nums {
		rightSum += val
	}
	leftSum := 0
	for idx, val := range nums {
		rightSum -= val
		ans[idx] = abs(leftSum - rightSum)
		leftSum += val
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findMiddleIndex(nums []int) int {
	rightSum := 0
	for _, v := range nums {
		rightSum += v
	}
	leftSum := 0
	for idx, v := range nums {
		rightSum -= v
		if leftSum == rightSum {
			return idx
		}
		leftSum += v
	}
	return -1
}
