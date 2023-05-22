package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ref : https://leetcode.com/problems/same-tree/solutions/560321/python-js-go-c-o-n-sol-by-dfs-w-visualization/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p != nil) && (q != nil) {
		// Both p and q are non-empty
		// check equality on both subtree
		return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		// At last one of them is empty
		// check whether both p and q are empty or not
		return p == q
	}
}

// ref : https://sweetcode.io/binary-search-tree-data-structure-in-go/
func (node *TreeNode) newNode(val int) {
	if val < node.Val {
		// if  left is nil go to the left
		if node.Left == nil {
			node.Left = &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   val,
			}
			// if left isn't nil - recursion here
		} else {
			node.Left.newNode(val)
		}
		// if the value is greater than the TreeNode , Go right
	} else {
		// if right is nil go to the left
		if node.Right == nil {
			node.Right = &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   val,
			}
		} else {
			// if right isn't nil - recursion here
			node.Right.newNode(val)
		}
	}
}
func main() {
	input := bufio.NewScanner(os.Stdin)
	list := []int{}
	pList, qList := []int{}, []int{}

	for input.Scan() {
		inputStr := input.Text()
		number, _ := strconv.Atoi(inputStr)
		if inputStr != "" {
			list = append(list, number)
		} else {
			break
		}
	}

	if len(list)%2 == 0 {
		for _, data := range list[0:3] {
			pList = append(pList, data)
			// p := newNode(data)
			// fmt.Println(p)
		}
		for _, data2 := range list[3:list[len(list)-1]] {
			qList = append(qList, data2)
			// q := newNode(data2)
			// fmt.Println("q :", q)
		}
	}
	fmt.Println(pList, qList)

}

/*
Input: p = [1,2,3], q = [1,2,3]
Output: true
*/
