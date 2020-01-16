package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent == nil {
		node = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else if node == node.Parent.Right {
		node.Parent.Right = rplc
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
func main() {
	root := &TreeNode{Data: "01"}
	BTreeInsertData(root, "07")
	BTreeInsertData(root, "12")
	BTreeInsertData(root, "10")
	BTreeInsertData(root, "05")
	BTreeInsertData(root, "02")
	BTreeInsertData(root, "03")

	node := BTreeSearchItem(root, "12")

	replacement := &TreeNode{Data: "55"}
	BTreeInsertData(replacement, "60")
	BTreeInsertData(replacement, "33")
	BTreeInsertData(replacement, "12")
	BTreeInsertData(replacement, "15")
	root = BTreeTransplant(root, node, replacement)
	buf := ""
	printTreeRight(root, &buf, "", "")
	fmt.Println(buf)

	fmt.Println()
	root1 := &TreeNode{Data: "33"}
	BTreeInsertData(root1, "5")
	BTreeInsertData(root1, "20")
	BTreeInsertData(root1, "52")
	BTreeInsertData(root1, "31")
	BTreeInsertData(root1, "13")
	BTreeInsertData(root1, "11")

	node1 := BTreeSearchItem(root1, "20")

	replacement1 := &TreeNode{Data: "55"}
	BTreeInsertData(replacement1, "60")
	BTreeInsertData(replacement1, "33")
	BTreeInsertData(replacement1, "12")
	BTreeInsertData(replacement1, "15")
	root1 = BTreeTransplant(root1, node1, replacement1)
	buf = ""
	printTreeRight(root1, &buf, "", "")
	fmt.Println(buf)

	fmt.Println()
	root2 := &TreeNode{Data: "03"}
	BTreeInsertData(root2, "39")
	BTreeInsertData(root2, "99")
	BTreeInsertData(root2, "11")
	BTreeInsertData(root2, "44")
	BTreeInsertData(root2, "14")
	BTreeInsertData(root2, "11")

	node2 := BTreeSearchItem(root2, "11")

	replacement2 := &TreeNode{Data: "55"}
	BTreeInsertData(replacement2, "60")
	BTreeInsertData(replacement2, "33")
	BTreeInsertData(replacement2, "12")
	BTreeInsertData(replacement2, "15")
	root1 = BTreeTransplant(root2, node2, replacement2)
	buf = ""
	printTreeRight(root1, &buf, "", "")
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
