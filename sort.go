package main

func QuickSelect(nums []int, k int) int {
	if k > len(nums) {
		panic("k cannot be greater than length of the input array")
	}
	return quickSelect(nums, k, 0, len(nums))
}

func quickSelect(nums []int, k, start, end int) int {
	if start == end {
		return nums[end]
	}
	p := Partition(nums, start, end)
	// k is 1 indexed, so subtracting one from makes it zero indexed
	if k-1 < p {
		return quickSelect(nums, k, start, p-1)
	} else if k-1 > p {
		return quickSelect(nums, k, p+1, end)
	} else {
		return nums[k-1]
	}
}

func Partition(nums []int, start, end int) int {
	// assuming last element in the range as the key
	key := nums[end]
	// all the element before key are <= key and all the element after key are > key
	// ptr is the boundary element arr[i] <= key < arr[j]
	// we are try to put key into its sorted index and return the index
	ptr := start
	for i := start; i <= end; i++ {
		if nums[i] <= key { // since are check for <= key this if statement will be executed at least once
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
	return ptr - 1 // Because of above assume this statement makes senses
}
