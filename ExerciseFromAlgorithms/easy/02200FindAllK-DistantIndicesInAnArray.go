package main

import "fmt"

func findKDistantIndices(nums []int, key int, k int) []int {
	return nil
}

// TODO
func main() {
	nums := []int{3, 4, 9, 1, 3, 9, 5}
	key, k := 9, 1
	fmt.Println(nums, key, k)
	list := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			list = append(list, i)
		}
	}
	for i, _ := range list {
		fmt.Println(nums[list[i]], list[i])
	}
}
