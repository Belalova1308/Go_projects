package main

import (
	"testing"
)

func TestAreToysBalancedParallel(t *testing.T) {
	t.Run("Test 1 in Parallel", func(t *testing.T) {
		t.Parallel()
		root := &TreeNode{HasToy: true}
		root.Left = &TreeNode{HasToy: true}
		root.Right = &TreeNode{HasToy: false}
		root.Right.Left = &TreeNode{HasToy: true}
		root.Right.Right = &TreeNode{HasToy: true}
		root.Left.Left = &TreeNode{HasToy: true}
		root.Left.Right = &TreeNode{HasToy: false}
		exp := [7]bool{true, true, false, true, true, false, true}
		result := unrollGarland(root)

		for i := 0; i < len(exp); i++ {
			if result[i] != exp[i] {
				t.Errorf("Result was incorrect, got: %t, want: %t.", result[i], exp[i])
			}

		}
	})
	t.Run("Test 2 in Parallel", func(t *testing.T) {
		t.Parallel()
		root := &TreeNode{HasToy: false}
		root.Left = &TreeNode{HasToy: false}
		root.Right = &TreeNode{HasToy: false}
		root.Right.Left = &TreeNode{HasToy: false}
		root.Right.Right = &TreeNode{HasToy: true}
		root.Right.Right.Right = &TreeNode{HasToy: true}
		root.Left.Left = &TreeNode{HasToy: true}
		root.Left.Right = &TreeNode{HasToy: true}
		exp := [8]bool{false, false, false, true, false, true, true, true}
		result := unrollGarland(root)

		for i := 0; i < len(exp); i++ {
			if result[i] != exp[i] {
				t.Errorf("Result was incorrect, got: %t, want: %t.", result[i], exp[i])
			}

		}
	})
}
