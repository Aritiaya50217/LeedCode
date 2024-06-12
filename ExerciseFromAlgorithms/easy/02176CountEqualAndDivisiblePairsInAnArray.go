package main

import "fmt"

func max(array []int, length int) int {
	for i := 0; i < length-1; i++ {
		if array[i] > array[i+1] {
			var temp = array[i]
			array[i] = array[i+1]
			array[i+1] = temp // swap
		}
	}
	var maxValue = array[length-1]

	return maxValue
}
func countPairs(nums []int, k int) int {
	length := len(nums)
	index := 0
	for i := 0; i < length; i++ {
		if nums[i] >= k {
			index = i
			return nums[index]
		}
	}
	return k

}

// TODO
func main() {
	nums := []int{3, 1, 2, 2, 2, 1, 3}
	k := 2
	res := countPairs(nums, k)
	fmt.Println(res)

	// length := len(nums)
	// index := 0
	// for i := 0; i < length; i++ {
	// 	if nums[i] >= k {
	// 		index = i
	// 		if nums[index] == nums[index+1] {
	// 			i *= i
	// 			fmt.Println(i)
	// 			return
	// 		}
	// 	}
	// 	fmt.Println(k)
	// 	return

	// }

}
