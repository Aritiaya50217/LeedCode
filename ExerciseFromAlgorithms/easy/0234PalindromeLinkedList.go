package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// ref : https://yuminlee2.medium.com/palindrome-string-and-linked-list-6e037ee5e270#4a9c
func isPalindrome(head *ListNode) bool {
	return false
}

// TODO
func main() {
	head := []int{1, 2, 2, 1}
	node := &ListNode{}
	for i := 0; i < len(head); i++ {
		// node1 := &ListNode{Val: 1}
		// node2 := &ListNode{Val: 2}
		// node1.Next = node2
		newNode := &ListNode{
			Val: head[i],
		}

		if node == nil {
			node = newNode
		} else {
			current := node
			for current.Next != nil {
				current = current.Next
			}
			current.Next = newNode
			fmt.Println(*current.Next)
		}
	}

}
