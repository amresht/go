package main 

/*
@author			amresht

@description	Linkedlist in Go : add, list, delete, insert operations
*/

import (
	"fmt"
	"math/rand"
)

type Node struct {
	info interface{}
	next *Node
}

type List struct{
	head *Node
}

func (myList *List) append (d interface{}) {
	// assign value to this node
	newNode := &Node{info: d, next:nil}
	// if list is empty create a new head
	if myList.head ==nil {
		myList.head = newNode
	} else {
		current := myList.head
		for current.next != nil{
			current  = current.next
		}
		current.next  = newNode
	}
} // end of method append


func printList (myList *List) {
	current := myList.head
	for current != nil {
		fmt.Printf ("-> %v", current.info)
		current  = current.next
	}
}


func main() {
	singleLinkedList := List{}
	for i =0; i<5;i++ {
		singleLinkedList.append(rand.Intn(100))
	}

	printList(&singleLinkedList)
}