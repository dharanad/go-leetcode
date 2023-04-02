package main

import (
	"container/list"
	"fmt"
)

func main() {
	//nums := []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}
	//fmt.Println(beautifulSubsets([]int{2, 4, 6}, 2))
	//fmt.Println(primeSubOperation([]int{4, 9, 6, 10}))
	fmt.Println(successfulPairs([]int{5, 13}, []int{1, 2, 3, 4, 5}, 7))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isValid(s string) bool {
	stack := list.New()
	openingBracketMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	// put opening into list
	// when closing comes check if its has its opening match in stack
	for _, c := range s {
		if c == '[' || c == '(' || c == '{' {
			stack.PushBack(c)
		} else {
			o := openingBracketMap[c]
			if stack.Len() != 0 && stack.Back().Value == o {
				stack.Remove(stack.Back())
			} else {
				return false
			}
		}
	}
	return stack.Len() == 0
}
