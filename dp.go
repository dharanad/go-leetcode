package main

import (
	"math"
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
	for idx, _ := range dp {
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
