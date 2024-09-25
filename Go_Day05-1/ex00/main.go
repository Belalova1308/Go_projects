package main

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func countToys(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.HasToy {
		count = 1
	}
	count += countToys(root.Left)
	count += countToys(root.Right)

	return count
}
func areToysBalanced(toy *TreeNode) bool {
	if toy == nil {
		return true
	}
	leftcount := countToys(toy.Left)
	rightcount := countToys(toy.Right)
	return leftcount == rightcount

}
