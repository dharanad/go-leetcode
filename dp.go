package main

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

}
