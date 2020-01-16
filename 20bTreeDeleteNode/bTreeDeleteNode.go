package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		del := BTreeMin(root.Right)

		root.Data = del.Data
		root.Right = BTreeDeleteNode(root.Right, del)
	}
	return root
}
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	}
	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}
func main() {
	root := &TreeNode{Data: "01"}
	BTreeInsertData(root, "07")
	BTreeInsertData(root, "12")
	BTreeInsertData(root, "05")
	BTreeInsertData(root, "10")
	BTreeInsertData(root, "02")
	BTreeInsertData(root, "03")

	node := BTreeSearchItem(root, "02")
	root = BTreeDeleteNode(root, node)
	buf := ""
	printTreeRight(root, &buf, "", "")
	fmt.Println(buf)

	root1 := &TreeNode{Data: "33"}
	BTreeInsertData(root1, "5")
	BTreeInsertData(root1, "52")
	BTreeInsertData(root1, "20")
	BTreeInsertData(root1, "31")
	BTreeInsertData(root1, "13")
	BTreeInsertData(root1, "11")

	node1 := BTreeSearchItem(root1, "20")
	root1 = BTreeDeleteNode(root1, node1)
	buf = ""
	printTreeRight(root1, &buf, "", "")
	fmt.Println(buf)

	root2 := &TreeNode{Data: "03"}
	BTreeInsertData(root2, "39")
	BTreeInsertData(root2, "99")
	BTreeInsertData(root2, "44")
	BTreeInsertData(root2, "11")
	BTreeInsertData(root2, "14")
	BTreeInsertData(root2, "11")

	node2 := BTreeSearchItem(root2, "03")
	root2 = BTreeDeleteNode(root2, node2)
	buf = ""
	printTreeRight(root2, &buf, "", "")
	fmt.Println(buf)

	root3 := &TreeNode{Data: "03"}
	BTreeInsertData(root3, "03")
	BTreeInsertData(root3, "94")
	BTreeInsertData(root3, "19")
	BTreeInsertData(root3, "24")
	BTreeInsertData(root3, "111")
	BTreeInsertData(root3, "01")

	node3 := BTreeSearchItem(root3, "03")
	root3 = BTreeDeleteNode(root3, node3)
	buf = ""
	printTreeRight(root3, &buf, "", "")
	fmt.Println(buf)
}
func printTreeRight(root *TreeNode, buffer *string, prefix string, childrenPrefix string) {

	if root != nil {
		*buffer += prefix
		pData := ""
		if root.Parent != nil {
			pData = root.Parent.Data
		} else {
			pData = "nil"
		}
		*buffer += fmt.Sprintf("%02s (%02s)", root.Data, pData)
		*buffer += "\n"
		if root.Right != nil {
			if root.Left != nil {
				printTreeRight(root.Right, buffer, childrenPrefix+"├── ", childrenPrefix+"│   ")
			} else {
				printTreeRight(root.Right, buffer, childrenPrefix+"└── ", childrenPrefix+"   ")
			}
		}
		if root.Left != nil {
			if root.Right != nil {
				printTreeRight(root.Left, buffer, childrenPrefix+"└── ", childrenPrefix+"    ")
			} else {
				printTreeRight(root.Left, buffer, childrenPrefix+"└── ", childrenPrefix+"    ")
			}
		}
	}
}
