package main

import "fmt"

func findFinalValue(nums []int, original int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == original {
			original *= 2
		}
	}
	return original
}

// TODO
func main() {
	// If original is found in nums, multiply it by two (i.e., set original = 2 * original).
	nums := []int{8, 19, 4, 2, 15, 3}
	original := 2
	res := findFinalValue(nums, original)
	fmt.Println(res)

}
