package main

import "fmt"

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		index := (nums[i] - 1) % len(nums)
		nums[index] = len(nums) + nums[index]
	}
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 2*len(nums) {
			nums[k] = i + 1
			k++
		}
	}
	list := nums[0:k]
	if len(list) == 0 {
		return false
	}

	return true
}

// TODO
func main() {
	nums := []int{3, 1}
	// res := containsDuplicate(nums)
	// fmt.Println(res)
	for i := 0; i < len(nums); i++ {
		index := (nums[i] - 1) % len(nums)
		nums[index] = len(nums) + nums[index]
	}
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 2*len(nums) {
			nums[k] = i + 1
			k++
		}
	}
	fmt.Println(nums)
	// checkList := nums[0:k]
	// if checkList == nil {
	// 	fmt.Println(false)
	// }
	// fmt.Println(checkList)

}
