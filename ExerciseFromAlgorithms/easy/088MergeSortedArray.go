package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// ref : https://blog.devgenius.io/leetcode-88-merge-sorted-array-get-solution-with-images-a6a40539c50
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	//Merge nums1 and nums2
	nums1 = append(nums1[:n], nums2...)
	// Sort nums1
	sort.Ints(nums1)
	return nums1

}

func main() {
	var m int
	var n int
	fmt.Scan(&m, &n)
	input := bufio.NewScanner(os.Stdin)
	list := []string{}
	nums1 := []int{}
	nums2 := []int{}
	for input.Scan() {
		inputStr := input.Text()
		if inputStr != "" {
			list = append(list, inputStr)
		} else {
			break
		}
	}
	for _, v := range list[:m] {
		number, _ := strconv.Atoi(v)
		nums1 = append(nums1, number)
	}
	for _, val := range list[len(list)-n:] {
		number, _ := strconv.Atoi(val)
		nums2 = append(nums2, number)
	}

	res := merge(nums1, m, nums2, n)
	fmt.Println(res)

}
