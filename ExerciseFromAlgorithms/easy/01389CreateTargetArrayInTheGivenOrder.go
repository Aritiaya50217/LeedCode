package main

import "fmt"

// res : https://leetcode.com/problems/create-target-array-in-the-given-order/solutions/2482692/go-simple-one-pass/
func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	for idx, i := range index {
		for j := len(target) - 1; j > i; j-- {
			target[j] = target[j-1]
		}
		target[i] = nums[idx]
	}
	return target
}
func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	res := createTargetArray(nums, index)
	fmt.Println(res)
}
