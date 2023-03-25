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

func countArrangement(n int) int {
	ans := 0
	taken := make([]bool, n+1)
	countArrangementHelper(1, taken, n, &ans)
	return ans
}

func countArrangementHelper(idx int, taken []bool, n int, ans *int) {
	if idx > n {
		*ans = *ans + 1
		return
	}
	// working on index idx
	for i := 1; i <= n; i++ {
		if !taken[i] && (i%idx == 0 || idx%i == 0) {
			taken[i] = true
			countArrangementHelper(idx+1, taken, n, ans)
			taken[i] = false
		}
	}
}

func countArrangementBitmask(n, mask, idx int) int {
	if idx == 0 {
		return 1
	}
	count := 0
	for i := 1; i <= n; i++ {
		if ((mask & (1 << i)) == 0) && (idx%i == 0 || i%idx == 0) {
			count += countArrangementBitmask(n, mask|(1<<i), idx-1)
		}
	}
	return count
}
