package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	for i := 0; bTreeApplyByLevel(root, i, f); i++ { //
	}
}

func bTreeApplyByLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) bool {
	if root == nil {
		return false
	}
	if level == 0 {
		f(root.Data)
		return true
	}
	left := bTreeApplyByLevel(root.Left, level-1, f)
	right := bTreeApplyByLevel(root.Right, level-1, f)
	return left || right
}
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	}
	if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func main() {
	root := &TreeNode{Data: "01"}
	BTreeInsertData(root, "07")
	BTreeInsertData(root, "12")
	BTreeInsertData(root, "05")
	BTreeInsertData(root, "10")
	BTreeInsertData(root, "02")
	BTreeInsertData(root, "03")
	BTreeApplyByLevel(root, fmt.Println)

	fmt.Println()
	root1 := &TreeNode{Data: "01"}
	BTreeInsertData(root1, "07")
	BTreeInsertData(root1, "12")
	BTreeInsertData(root1, "05")
	BTreeInsertData(root1, "10")
	BTreeInsertData(root1, "02")
	BTreeInsertData(root1, "03")
	BTreeApplyByLevel(root1, fmt.Print)
	fmt.Println()
}
