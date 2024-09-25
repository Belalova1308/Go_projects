package main

import "testing"

func TestAreToysBalancedParallel(t *testing.T) {
	t.Run("Test 1 in Parallel", func(t *testing.T) {
		t.Parallel()
		root := &TreeNode{HasToy: false}
		root.Left = &TreeNode{HasToy: false}
		root.Right = &TreeNode{HasToy: true}
		root.Left.Left = &TreeNode{HasToy: false}
		root.Left.Right = &TreeNode{HasToy: true}
		result := areToysBalanced(root)
		if result != true {
			t.Errorf("Result was incorrect, got: %t, want: %t.", result, true)
		}
	})
	t.Run("Test 2 in Parallel", func(t *testing.T) {
		t.Parallel()
		root := &TreeNode{HasToy: true}
		root.Left = &TreeNode{HasToy: true}
		root.Right = &TreeNode{HasToy: false}
		result := areToysBalanced(root)
		if result != false {
			t.Errorf("Result was incorrect, got: %t, want: %t.", result, false)
		}
	})
}
