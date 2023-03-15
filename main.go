package main

import "fmt"

func main() {
	fmt.Println("Hello, leetcode")
	fmt.Println(coloredCells(1))
	fmt.Println(coloredCells(2))
	fmt.Println(coloredCells(3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
