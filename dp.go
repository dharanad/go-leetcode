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
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, val := range coins {
			if i >= val {
				dp[i] = Min(dp[i], dp[i-val])
			}
		}
		if dp[i] != math.MaxInt {
			dp[i]++
		}
	}
	return dp[amount]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
