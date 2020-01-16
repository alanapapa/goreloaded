package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	cur := l.Head
	prev := l.Head
	for cur != nil && l.Head.Data == data_ref {
		cur = l.Head.Next
		l.Head = cur
	}
	for cur != nil {
		if cur.Data == data_ref {
			cur = cur.Next
			prev.Next = cur
		} else {
			prev = cur
			cur = cur.Next
		}
	}
}
func ListPushBack(l *List, data interface{}) {
	a := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = a
	} else {
		l.Tail.Next = a
	}
	l.Tail = a
}
func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}
func main() {

	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link, 1)
	ListPushBack(link, "Hello")
	ListPushBack(link, 1)
	ListPushBack(link, "There")
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, "How")
	ListPushBack(link, 1)
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}
