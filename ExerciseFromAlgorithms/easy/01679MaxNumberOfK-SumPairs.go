package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// res : https://leetcode.com/problems/max-number-of-k-sum-pairs/solutions/2006150/go-best-running-time-solution-100-faster-time-101ms-100-less-mem-usage/
func maxOperations(nums []int, k int) int {
	hashmap := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		// base case
		if len(nums) <= 1 {
			return ans
		}
		// hashmap will store values in the form
		if nums[i] >= k {
			continue
		}
		// check if counterpart of nums[i] is present
		counterpart := k - nums[i]
		if val, exists := hashmap[counterpart]; exists {
			if val != 0 {
				ans++
			}
			// taking care of duplicates
			if val-1 == 0 {
				delete(hashmap, counterpart)
			} else {
				hashmap[counterpart]--
			}
		} else {
			hashmap[nums[i]]++
		}
	}

	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	nums := []int{}

	for input.Scan() {
		inputStr := input.Text()
		numbers, _ := strconv.Atoi(inputStr)
		if inputStr != "" {
			nums = append(nums, numbers)
		} else {
			break
		}
	}
	last := nums[len(nums)-1:]
	k := last[0]
	list := nums[:len(nums)-1]
	res := maxOperations(list, k)
	fmt.Println(res)

}

// nums = [1,2,3,4], k = 5
