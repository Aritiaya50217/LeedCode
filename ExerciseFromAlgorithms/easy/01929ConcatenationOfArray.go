package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}

// TODO
func main() {
	input := bufio.NewScanner(os.Stdin)
	list := []int{}
	for input.Scan() {
		inputStr := input.Text()
		number, _ := strconv.Atoi(inputStr)
		if inputStr != "" {
			list = append(list, number)
		} else {
			break
		}
	}
	res := getConcatenation(list)
	fmt.Println("res :", res)

}
