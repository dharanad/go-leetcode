package main

import "fmt"

func main() {
	nums := []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}
	fmt.Println(maxNumOfMarkedIndices(nums))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
