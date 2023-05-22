package main

import "fmt"

func findDuplicates(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		index := (nums[i] - 1) % length
		nums[index] = length + nums[index]
	}
	k := 0
	for i := 0; i < length; i++ {
		if nums[i] > 2*length {
			nums[k] = i + 1
			k++
		}
	}
	return nums[0:k]
}

func findLucky(arr []int) int {
	list := findDuplicates(arr)
	res := 0
	for _, v := range list {
		res = v
	}
	return res

	// list := findDuplicates(arr)
	// total := 0
	// for i, v := range list {
	// 	if len(list) >= 2 {
	// 		var next = list[i+1]
	// 		total = list[i] - next
	// 		return total
	// 	} else {
	// 		total = v
	// 	}
	// }
	// return total
}

func main() {
	input := []int{2, 2, 2, 3, 3}
	res := findLucky(input)
	fmt.Println(res)
}
