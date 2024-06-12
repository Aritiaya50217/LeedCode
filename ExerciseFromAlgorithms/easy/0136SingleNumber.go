package main

import "fmt"

func singleNumber(nums []int) int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		exists := false
		for v := 0; v < i; v++ {
			if nums[v] != nums[i] {
				exists = true
				// break
			}
		}
		if !exists {
			result = append(result, nums[i])
		}
	}
	return result[0]
}

// TODO
func main() {
	nums := []int{2, 2, 1}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		exists := false
		for v := 0; v < i; v++ {
			if nums[v] == nums[i] {
				exists = true
				// break
			}
		}
		if !exists {
			result = append(result, nums[i])
		}
	}
	fmt.Println(result)
}
