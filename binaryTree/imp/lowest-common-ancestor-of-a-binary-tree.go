package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LCA(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// lc and rc will be nil if root children does not contain p and q.
	lc := LCA(root.Left, p, q)
	rc := LCA(root.Right, p, q)

	if root == p || root == q {
		return root
	}

	if lc == nil {
		return rc
	} else if rc == nil {
		return lc
	} else {
		return root
	}
}
