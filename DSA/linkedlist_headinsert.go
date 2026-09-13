package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func insert(head *node, data int) *node {
	//n := &node{data: data}
	n := &node{data, nil}
	n.next = head
	return n
	/*
		if head == nil {
			return n
		} else {
			n.next = head
			return n
		}*/
}
func printList(head *node) {
	for head != nil {
		fmt.Println(head.data, " -> ")
		head = head.next
	}
	fmt.Println(nil)
}
func main() {
	var link *node
	link = insert(link, 1)
	link = insert(link, 3)
	link = insert(link, 5)
	printList(link)
}
