package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a := BTreeLevelCount(root.Left)
	b := BTreeLevelCount(root.Right)
	if b > a {
		return b + 1
	}
	return a + 1
}
