package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	head := &ListNode{0, nil}
	curr := head
	for {
		if list1 == nil || list2 == nil {
			break
		}
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return head.Next
}

func main() {
	// input := bufio.NewScanner(os.Stdin)
	// list := []string{}
	// // numberList := []int{}
	// for input.Scan() {
	// 	inputStr := input.Text()
	// 	list = append(list, inputStr)
	// 	if inputStr == "" {
	// 		break
	// 	}
	// }
	// fmt.Println(list)

	test := []ListNode{{1, nil}, {2, nil}, {3, nil}}
	for _, val := range test {
		fmt.Println(val.Val)
	}

}
