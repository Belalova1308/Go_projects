package main

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func levelOrder(root *TreeNode) [][]bool {
	m := make(map[int][]bool)
	foo(root, 0, m)
	s := make([][]bool, len(m))
	for k, v := range m {
		s[k] = v
	}
	return s
}

func foo(n *TreeNode, level int, m map[int][]bool) {
	if n == nil {
		return
	}
	m[level] = append(m[level], n.HasToy)
	foo(n.Left, level+1, m)
	foo(n.Right, level+1, m)
}

func unrollGarland(root *TreeNode) []bool {
	levelArray := levelOrder(root)
	var result []bool
	for k, v := range levelArray {
		if k%2 == 0 {
			for i := len(v) - 1; i >= 0; i-- {
				result = append(result, v[i])
			}
		} else {
			for i := 0; i < len(v); i++ {
				result = append(result, v[i])
			}
		}
	}
	return result
}
