package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func removeDuplicates(nums []int) []int {
	val := map[int]bool{}
	result := []int{}
	for i := range nums {
		if val[nums[i]] != true {
			val[nums[i]] = true
			result = append(result, nums[i])
		}
	}
	//length := len(nums)
	// l := length - len(result)
	// for i := 1; i <= l; i++ {
	// 	result = append(result, "-")
	// }

	return result
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	list := []int{}
	for input.Scan() {
		inputStr := input.Text()
		if inputStr == "" {
			break
		} else {
			number, _ := strconv.Atoi(inputStr)
			list = append(list, number)
		}
	}
	res := removeDuplicates(list)
	fmt.Println(res)
}
