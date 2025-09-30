package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (head *LinkedList) Insert(value int) {
	if head != nil {
		head.Next = &LinkedList{Value: value}
	}
}

func printLinkedList(head *LinkedList) {
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func main() {
	var node1 *LinkedList = &LinkedList{1, nil}
	node1.Next = &LinkedList{2, nil}
	node1.Next.Next = &LinkedList{3, nil}
	node1.Next.Next.Insert(4)

	printLinkedList(node1)
}
