package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (ll *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data, ll.head}
	ll.head = newNode
}

func (ll *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data, nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) DeleteNode(key int) {
	current := ll.head
	var prev *Node
	if current != nil && current.data == key {
		ll.head = current.next
		return
	}
	for current != nil && current.data != key {
		prev = current
		current = current.next
	}

	// If the key was not present
	if current == nil {
		fmt.Println("Key not found:", key)
		return
	}

	// Unlink the node from the list
	prev.next = current.next
}

func (ll *LinkedList) Search(key int) bool {
	current := ll.head
	for current != nil {
		if current.data == key {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList) Display() {
	current := ll.head
	if current == nil {
		fmt.Println("Linked List is empty")
		return
	}
	fmt.Print("Linked List: ")
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func print_ll() {
	ll := LinkedList{}

	// Insert elements
	ll.InsertAtEnd(10)
	ll.InsertAtEnd(20)
	ll.InsertAtBeginning(5)
	ll.InsertAtEnd(30)

	// Display the list
	ll.Display() // Output: Linked List: 5 -> 10 -> 20 -> 30 -> nil

	ll.DeleteNode(20)
	ll.Display() // Output: Linked List: 5 -> 10 -> 30 -> nil

	// Search for elements
	fmt.Println("Search for 20:", ll.Search(20)) // Output: true
	fmt.Println("Search for 40:", ll.Search(40)) // Output: false

}
