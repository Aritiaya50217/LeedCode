package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//  Definition for singly-linked list.
type ListNode struct {
	Val      int
	Next     *ListNode
	Previous *ListNode
}

var root = new(ListNode)

// add value in ListNode
func addNode(t *ListNode, v int) int {
	// check root
	if root == nil {
		t = &ListNode{v, nil, nil}
		root = t
		return 0
	}
	if v == t.Val {
		fmt.Println("Node already exists: ", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &ListNode{v, nil, nil}
		return -2
	}
	return addNode(t.Next, v)
}

// check empty
func traverse(t *ListNode) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Val)
		t = t.Next
	}
	fmt.Println()
}

func deleteDuplicates(head *ListNode) *ListNode {
	// check head
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	var t, test = &ListNode{}, &ListNode{}
	for input.Scan() {
		inputStr := input.Text()
		number, _ := strconv.Atoi(inputStr)
		if inputStr != "" {
			t = &ListNode{
				Val:  number,
				Next: nil,
			}
			test = deleteDuplicates(t)
		} else {
			break
		}
	}
	fmt.Println(test.Val)
}
