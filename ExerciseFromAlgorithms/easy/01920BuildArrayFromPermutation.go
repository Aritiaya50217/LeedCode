package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ref : https://medium.com/theleanprogrammer/build-array-from-permutation-907ff395eb40
/*
	Input: nums = [0,2,1,5,3,4] , n = 6  => len(nums)
	Output: [0,1,2,4,5,3]
	Explanation: The array ans is built as follows:
	// [nums[index]]
	ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
    	= [nums[0], nums[2], nums[1], nums[5], nums[3], nums[4]]  // value in nums[]
    	= [0,1,2,4,5,3]
*/

/*

Algorithms :

	length := len(nums)
	for i := 0; i < len(nums); i++ {
		// ex []int{1,2,3} => nums[nums[0]] = nums[1] => 2
		res[i] = nums[nums[i]]
		test := length*(nums[nums[i]]) + nums[i]
		fmt.Printf("%v ", test)  0 8 13 29 33 22
	}

ex : 0,2,1,5,3,4
nums
	0  2  1  5  3  4
res
	0  1  2  4  5  3

	nums[0] = res[res[0]] = res[0] = index 0
	nums[1] = res[res[1]] = res[2] = index 1
	nums[2] = res[res[2]] = res[1] = index 2
	nums[3] = res[res[3]] = res[5] = index 4
	nums[4] = res[res[4]] = res[3] = index 5
	nums[5] = res[res[5]] = res[4] = index 3


*/

func buildArray(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// ex []int{1,2,3} => nums[nums[0]] = nums[1] => 2
		result[i] = nums[nums[i]]
	}
	return result
}
func main() {
	input := bufio.NewScanner(os.Stdin)
	nums := []int{}
	for input.Scan() {
		inputStr := input.Text()
		number, _ := strconv.Atoi(inputStr)
		if inputStr != "" {
			nums = append(nums, number)
		} else {
			break
		}
	}

	res := []int{}
	res = buildArray(nums)
	fmt.Printf("%v ", res)
}
