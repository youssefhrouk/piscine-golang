package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}


func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		}
		temp := BTreeMin(root.Right)

		root.Data = temp.Data
		root.Right = BTreeDeleteNode(root.Right, temp)
	}
	return root
}