package main

import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	listPushBack(l, data_ref)
	for i := l; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
		}
	}
	return l
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

	var link *NodeI

	link = listPushBack(link, 0)

	link = SortListInsert(link, 39)
	PrintList(link)

	var link1 *NodeI

	link1 = listPushBack(link1, 0)
	link1 = listPushBack(link1, 1)
	link1 = listPushBack(link1, 2)
	link1 = listPushBack(link1, 3)
	link1 = listPushBack(link1, 4)
	link1 = listPushBack(link1, 5)
	link1 = listPushBack(link1, 24)
	link1 = listPushBack(link1, 25)
	link1 = listPushBack(link1, 54)

	link1 = SortListInsert(link1, 33)
	PrintList(link1)

	var link2 *NodeI

	link2 = listPushBack(link2, 0)
	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 18)
	link2 = listPushBack(link2, 33)
	link2 = listPushBack(link2, 37)
	link2 = listPushBack(link2, 37)
	link2 = listPushBack(link2, 39)
	link2 = listPushBack(link2, 52)
	link2 = listPushBack(link2, 53)
	link2 = listPushBack(link2, 57)

	link2 = SortListInsert(link2, 53)
	PrintList(link2)

	var link3 *NodeI

	link3 = listPushBack(link3, 0)
	link3 = listPushBack(link3, 5)
	link3 = listPushBack(link3, 18)
	link3 = listPushBack(link3, 24)
	link3 = listPushBack(link3, 28)
	link3 = listPushBack(link3, 35)
	link3 = listPushBack(link3, 42)
	link3 = listPushBack(link3, 45)
	link3 = listPushBack(link3, 52)

	link3 = SortListInsert(link3, 52)
	PrintList(link3)

	var link4 *NodeI

	link4 = listPushBack(link4, 0)
	link4 = listPushBack(link4, 12)
	link4 = listPushBack(link4, 20)
	link4 = listPushBack(link4, 23)
	link4 = listPushBack(link4, 23)
	link4 = listPushBack(link4, 24)
	link4 = listPushBack(link4, 30)
	link4 = listPushBack(link4, 41)
	link4 = listPushBack(link4, 53)
	link4 = listPushBack(link4, 57)
	link4 = listPushBack(link4, 59)

	link4 = SortListInsert(link4, 38)
	PrintList(link4)
}
