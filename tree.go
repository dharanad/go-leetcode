package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	/*

	   Pre Order
	   D L R
	   3 9 20 15 7

	   In Order
	   L D R
	   9 3 15 20 7

	   What goes in left and right can decide based on Inrder Order
	*/
	idx := 0
	return BuildTreeHelper(preorder, inorder, &idx)
}

func BuildTreeHelper(preorder, inorder []int, idx *int) *TreeNode {
	if *idx == len(preorder) || 0 == len(inorder) {
		return nil
	}
	node := &TreeNode{Val: preorder[*idx]}
	*idx = *idx + 1
	var p = InorderIndex(inorder, node.Val)
	if p != -1 {
		node.Left = BuildTreeHelper(preorder, inorder[:p], idx)
		node.Right = BuildTreeHelper(preorder, inorder[p+1:], idx)
	}
	return node
}

func InorderIndex(inorder []int, val int) int {
	for i, v := range inorder {
		if val == v {
			return i
		}
	}
	return -1
}
