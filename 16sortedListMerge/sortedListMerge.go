package main

import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	} else {
		if n1.Data < n2.Data {
			n1.Next = SortedListMerge(n1.Next, n2)
			return n1
		}

		n2.Next = SortedListMerge(n1, n2.Next)
		return n2
	}
}
func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link0 *NodeI
	var link02 *NodeI

	PrintList(SortedListMerge(link02, link0))

	var link1 *NodeI
	var link12 *NodeI

	link12 = listPushBack(link12, 4)
	link12 = listPushBack(link12, 2)
	link12 = listPushBack(link12, 5)
	link12 = listPushBack(link12, 9)
	link12 = listPushBack(link12, 1)
	link12 = listPushBack(link12, 12)
	link12 = listPushBack(link12, 30)
	link12 = listPushBack(link12, 20)

	link1 = listPushBack(link1, 4)
	link1 = listPushBack(link1, 1)
	link1 = listPushBack(link1, 5)
	link1 = listPushBack(link1, 20)

	PrintList(SortedListMerge(link12, link1))

	// var link2 *NodeI
	// var link22 *NodeI

	// link2 = listPushBack(link2, 4)
	// link2 = listPushBack(link2, 4)
	// link2 = listPushBack(link2, 6)
	// link2 = listPushBack(link2, 9)
	// link2 = listPushBack(link2, 13)
	// link2 = listPushBack(link2, 18)
	// link2 = listPushBack(link2, 20)
	// link2 = listPushBack(link2, 20)

	// PrintList(SortedListMerge(link2, link22))

	// var link3 *NodeI
	// var link32 *NodeI

	// link3 = listPushBack(link3, 0)
	// link3 = listPushBack(link3, 7)
	// link3 = listPushBack(link3, 39)
	// link3 = listPushBack(link3, 92)
	// link3 = listPushBack(link3, 97)
	// link3 = listPushBack(link3, 93)
	// link3 = listPushBack(link3, 91)
	// link3 = listPushBack(link3, 28)
	// link3 = listPushBack(link3, 64)

	// link32 = listPushBack(link32, 80)
	// link32 = listPushBack(link32, 23)
	// link32 = listPushBack(link32, 27)
	// link32 = listPushBack(link32, 30)
	// link32 = listPushBack(link32, 85)
	// link32 = listPushBack(link32, 81)
	// link32 = listPushBack(link32, 75)
	// link32 = listPushBack(link32, 70)

	// PrintList(SortedListMerge(link32, link3))

	// var link4 *NodeI
	// var link42 *NodeI

	// link4 = listPushBack(link4, 0)
	// link4 = listPushBack(link4, 2)
	// link4 = listPushBack(link4, 11)
	// link4 = listPushBack(link4, 30)
	// link4 = listPushBack(link4, 54)
	// link4 = listPushBack(link4, 56)
	// link4 = listPushBack(link4, 70)
	// link4 = listPushBack(link4, 79)
	// link4 = listPushBack(link4, 99)

	// link42 = listPushBack(link42, 23)
	// link42 = listPushBack(link42, 28)
	// link42 = listPushBack(link42, 38)
	// link42 = listPushBack(link42, 67)
	// link42 = listPushBack(link42, 67)
	// link42 = listPushBack(link42, 79)
	// link42 = listPushBack(link42, 95)
	// link42 = listPushBack(link42, 97)

	// PrintList(SortedListMerge(link42, link4))

	// var link5 *NodeI
	// var link52 *NodeI

	// link5 = listPushBack(link5, 0)
	// link5 = listPushBack(link5, 3)
	// link5 = listPushBack(link5, 8)
	// link5 = listPushBack(link5, 8)
	// link5 = listPushBack(link5, 13)
	// link5 = listPushBack(link5, 19)
	// link5 = listPushBack(link5, 34)
	// link5 = listPushBack(link5, 38)
	// link5 = listPushBack(link5, 46)

	// link52 = listPushBack(link52, 7)
	// link52 = listPushBack(link52, 39)
	// link52 = listPushBack(link52, 45)
	// link52 = listPushBack(link52, 53)
	// link52 = listPushBack(link52, 59)
	// link52 = listPushBack(link52, 70)
	// link52 = listPushBack(link52, 76)
	// link52 = listPushBack(link52, 79)

	// PrintList(SortedListMerge(link52, link5))
}
