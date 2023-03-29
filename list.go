package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeSet := make(map[*ListNode]struct{})
	ptr := headA
	for ptr != nil {
		nodeSet[ptr] = struct{}{}
		ptr = ptr.Next
	}
	ptr = headB
	for ptr != nil {
		if _, ok := nodeSet[ptr]; ok {
			return ptr
		}
		ptr = ptr.Next
	}
	return nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		9 -> 9 -> 9 -> 9
		1
		0 0 0 0
		rem = 1
	*/
	head := &ListNode{}
	ptr := head
	rem := 0
	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		sum += rem

		ptr.Next = &ListNode{Val: sum % 10}

		rem = sum / 10
		ptr = ptr.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if rem > 0 {
		ptr.Next = &ListNode{Val: rem}
	}
	return head.Next
}

func removeElement(nums []int, val int) int {
	ptr := 0
	for idx, v := range nums {
		if v != val {
			nums[idx], nums[ptr] = nums[ptr], nums[idx]
			ptr++
		}
	}
	return ptr
}

// [1,1,2,2,3,4]
func removeDuplicates(nums []int) int {
	ptr := 0
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == nums[ptr] {
			continue
		} else {
			ptr++
			nums[ptr] = nums[idx]
		}
	}
	return ptr + 1
}
