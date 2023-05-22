package main

import (
	"bufio"
	"fmt"
	"os"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	list := []int{}
	return list
}

func traverse(t *TreeNode) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Println(t.Val, " ")
	traverse(t.Right)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	list := []string{}
	for input.Scan() {
		inputStr := input.Text()
		if inputStr != "" {
			list = append(list, inputStr)
		} else {
			break
		}
	}
	fmt.Println(list)

}
