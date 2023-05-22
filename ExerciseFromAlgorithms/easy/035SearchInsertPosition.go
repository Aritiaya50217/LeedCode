package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func searchInsert(nums []int, target int) int {
	lenNums := len(nums)
	index := -1
	if target <= nums[0] {
		return 0
	}
	if target > nums[lenNums-1] {
		return lenNums
	}
	for i := 0; i < lenNums; i++ {
		if target <= nums[i] {
			index = i
			break
		}
	}
	return index
}
func main() {
	input := bufio.NewScanner(os.Stdin)
	nums := []int{}
	target := 0
	for input.Scan() {
		inputStr := input.Text()
		number, _ := strconv.Atoi(inputStr)
		if inputStr != "" {
			nums = append(nums, number)
		} else {
			break
		}
	}

	last := nums[len(nums)-1:]
	target = last[0]
	// for _, v := range nums[:len(nums)-1] {
	// 	if target != v {
	// 		nums = append(nums[:len(nums)-1], target)
	// 		// sort int
	// ref : https://golangbyexample.com/search-insert-position-golang/
	// 		sort.Slice(nums[:], func(i, j int) bool {
	// 			return (nums[:][i] < nums[:][j])
	// 		})
	// 	}
	// }
	// fmt.Println("----------")
	// for i := 0; i < len(nums); i++ {
	// 	if target == nums[i] {
	// 		fmt.Println(i)
	// 		return
	// 	}
	// }

	fmt.Println("--- Use Algorithms --- ")
	res := searchInsert(nums, target)
	fmt.Println(res)

}
