package main

import (
	"fmt"
	"strconv"
	"strings"
)

func runningSum(nums []int) []int {
	list := []int{}
	for _, v := range nums {
		list = append(list, v)
	}

	return list
}

func main() {
	nums := []int{1, 2, 3, 4}
	list := runningSum(nums)
	sumList := []int{}
	var k int
	for i, val := range list {
		for space := 1; space <= val-i; space++ {
			k += val
			sumList = append(sumList, k)
		}
	}
	resList := []string{}
	for _, val := range sumList {
		numbers := strconv.Itoa(val)
		resList = append(resList, numbers)
	}
	res := strings.Join(resList, ",")
	fmt.Printf("[%v]", res)
}
