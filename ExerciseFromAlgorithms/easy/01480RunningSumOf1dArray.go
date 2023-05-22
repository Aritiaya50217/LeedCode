package main

import "fmt"

func runningSum(nums []int) []int {
	list := []int{}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		list = append(list, sum)
	}
	return list
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := runningSum(nums)
	fmt.Println(res)
}
