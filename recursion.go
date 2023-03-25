package main

func beautifulSubsets(nums []int, k int) int {
	ans := 0
	numCountMap := make(map[int]int)
	beautifulSubsetsHelper(nums, k, 0, numCountMap, &ans)
	return ans - 1
}

func beautifulSubsetsHelper(nums []int, k, idx int, countMap map[int]int, ans *int) {
	if idx == len(nums) {
		*ans = *ans + 1
		return
	}
	// not include the current number
	beautifulSubsetsHelper(nums, k, idx+1, countMap, ans)

	// check if number at index is valid
	if i, j := countMap[nums[idx]+k], countMap[nums[idx]-k]; i == 0 && j == 0 {
		// include into the subset
		countMap[nums[idx]]++
		beautifulSubsetsHelper(nums, k, idx+1, countMap, ans)
		// after backtracking discard from subset
		countMap[nums[idx]]--
	}
}
