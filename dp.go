package main

import (
	"math"
	"sort"
)

/*
Knapsack means a bag. Just like a school or a travel or laptop back
In this problem either we can take everything or take nothing from a bag

Variants
1. 0/1 Knapsack
2. Bounded Knapsack
3. Unbounded


Each item in the bag has a value and cost and associated to it
We to need to maximize the value by not exceeding the weight or with the exact weight

*/

func ZeroOneKnapsack(cost, value []int, idx, w, n int) int {
	/*
		at a particular index i can pick a item or not pick a item
		if we pick item at idx, we increase the profit by value at idx
		and decrease the weight by weight at i
		and compute the same sub problem for rest of the array
	*/
	if idx == n || w == 0 { // if there is no array left or weight is zero
		return 0
	}
	if cost[idx] <= w {
		// max of picking and not picking the item
		return MaxInt(
			value[idx]+ZeroOneKnapsack(cost, value, idx+1, w-cost[idx], n),
			ZeroOneKnapsack(cost, value, idx+1, w, n))
	} else {
		return ZeroOneKnapsack(cost, value, idx+1, w, n)
	}
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func FractionalKnapsack(cost, value []int, idx, w, n int) int {
	return -1
}

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	// dp[i] -> no of way to reach to reach top
	// dp[i] -> dp[i-1] + dp[i-2] by climbing one or two steps
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func coinChange(coins []int, amount int) int {
	if x := _coinChange(coins, amount); x == math.MaxInt {
		return -1
	} else {
		return x
	}
}

func _coinChange(coins []int, amount int) int {
	/*
		dp[i] -> min no of coins required to make amount i
	*/
	if amount == 0 {
		return 0
	}
	res := math.MaxInt
	for _, val := range coins {
		if amount >= val {
			res = Min(res, _coinChange(coins, amount-val))
		}
	}
	if res != math.MaxInt {
		return 1 + res
	}
	return res
}

func __coinChange(coins []int, amount int) int {
	// ideally the amount of coins required should not exceed the total amount of coins
	// so we can assume amount+1 as maxInt in this context
	maxInt := amount + 1
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = maxInt
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, val := range coins {
			if i >= val {
				dp[i] = Min(dp[i], 1+dp[i-val])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	// sub seq is subset that may contain non-contiguous elements
	// each element in the array is a seq of length 1
	// dp[i] -> length of longest inc sub-sequence ending with element nums[i]
	dp := make([]int, len(nums))
	res := 0
	for idx := range dp {
		dp[idx] = 1
		maxSubSeqIdx := idx
		for j := idx - 1; j >= 0; j -= 1 {
			if nums[idx] > nums[j] && dp[j] >= dp[maxSubSeqIdx] {
				maxSubSeqIdx = j
			}
		}
		if maxSubSeqIdx != idx {
			dp[idx] += dp[maxSubSeqIdx]
		}
		// coz sub seq can end at any index and we need to find the max LIS
		res = Max(res, dp[idx])
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _lengthOfLIS(nums []int) int {
	/*
		2,3,7,101,3,7,101,18

	*/
	arrLen := len(nums)
	length := 1
	for idx := 1; idx < arrLen; idx++ {
		if nums[idx] > nums[length] {
			length++
		} else { // so nums[idx] would replace some or the other element
			// because of the invariant that [0..length-1] is sort and inc array
			lbIdx := lowerBound(nums[idx], 0, length-1, nums)
			nums[lbIdx] = nums[idx]
		}
	}
	return length
}

func lowerBound(x, lo, hi int, nums []int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2 // left biased mid, since we are moving away from lo
		if nums[mid] >= x {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if nums[lo] >= x {
		return lo
	}
	return -1
}

func maxEnvelopes(envelopes [][]int) int {
	/*
	   fits iff h & w are >
	*/
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	arrLen := len(envelopes)
	length := 1
	for i := 1; i < arrLen; i++ {
		if gt(envelopes, i, length-1) {
			envelopes[length] = envelopes[i]
			length++
		} else {
			lo := 0
			hi := length - 1
			for lo < hi {
				mid := lo + (hi-lo)/2
				if gte(envelopes, mid, i) {
					hi = mid
				} else {
					lo = mid + 1
				}
			}
			envelopes[lo] = envelopes[i]
		}
	}
	return length
}

// gt check if first envelop is greater than the second one
func gt(envelopes [][]int, first, second int) bool {
	return envelopes[second][0] < envelopes[first][0] && envelopes[second][1] < envelopes[first][1]
}

func gte(envelopes [][]int, first, second int) bool {
	return envelopes[second][0] <= envelopes[first][0] && envelopes[second][1] <= envelopes[first][1]
}

func subarraySum(nums []int, k int) int {
	countMap := make(map[int]int)
	res := 0
	sum := 0
	for _, val := range nums {
		sum += val
		if sum == k {
			res++
		}
		if v, ok := countMap[sum-k]; ok {
			res += v
		}
		countMap[sum] += 1
	}
	return res
}

func repairCars(ranks []int, cars int) int64 {
	lo, hi := int64(0), int64(100*1e5*1e5)
	for lo < hi {
		var mid = lo + (hi-lo)/2
		if canRepair(ranks, cars, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canRepair(ranks []int, cars int, time int64) bool {
	var carsRepaired int64 = 0
	for _, val := range ranks {
		carsRepaired += int64(math.Sqrt(float64(time / int64(val))))
	}
	return carsRepaired >= int64(cars)
}

func evenOddBit(n int) []int {
	even, odd := 0, 0
	for i := 0; i < 31 && n > 0; i++ {
		if n&1 == 1 {
			if i&1 == 0 {
				even++
			} else {
				odd++
			}
		}
		n = n >> 1
	}
	return []int{even, odd}
}

func checkValidGrid(grid [][]int) bool {
	dx := []int{1, -1, 1, -1, -2, -2, 2, 2}
	dy := []int{-2, -2, 2, 2, 1, -1, 1, -1}
	counter := 0
	n := len(grid)
	i, j := 0, 0
	for true {
		for k := 0; k < 8; k++ {
			ni := i + dx[k]
			nj := j + dy[k]
			if ni < 0 || ni >= n || nj < 0 || nj >= n {
				continue
			}
			if grid[ni][nj] == counter+1 {
				i = ni
				j = nj
				break
			}
		}
		if grid[i][j] != counter+1 {
			break
		}
		counter++
	}
	return counter == (n*n - 1)
}

func rob(nums []int) int {
	// dp[i] -> max amount of robbed till index i
	arrLen := len(nums)
	if arrLen == 1 {
		return nums[0]
	}
	if arrLen == 2 {
		return MaxInt(nums[0], nums[1])
	}
	dp := make([]int, arrLen)
	dp[0] = nums[0]
	dp[1] = nums[1]
	for i := 2; i < arrLen; i++ {
		dp[i] = MaxInt(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[arrLen-1]
}

func canPartitionKSubsets(nums []int, k int) bool {
	arrLen := len(nums)
	totalArraySum := 0
	for _, val := range nums {
		totalArraySum += val
	}
	if totalArraySum%k != 0 {
		return false
	}

	targetSum := totalArraySum / k
	subsetCount := 1 << arrLen
	dp := make([]int, subsetCount)
	for idx := range dp {
		dp[idx] = -1
	}
	dp[0] = 0
	for mask := 0; mask < subsetCount; mask++ {
		if dp[mask] == -1 {
			continue
		}

		for j := 0; j < arrLen; j++ {
			if ((mask & (1 << j)) == 0) && dp[mask]+nums[j] <= targetSum {
				dp[mask|(1<<j)] = (dp[mask] + nums[j]) % targetSum
			}
		}
	}
	return dp[subsetCount-1] == 0
}
