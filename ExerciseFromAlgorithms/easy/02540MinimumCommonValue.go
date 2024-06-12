package main

import "fmt"

func min(array []int, length int) int {
	minIndex := 0
	for i := 1; i < length; i++ {
		if array[minIndex] > array[i] {
			minIndex = i
		}
	}
	return array[minIndex]
}
func getCommon(nums1 []int, nums2 []int) int {
	list := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				list = append(list, nums1[i])
			}
		}
	}
	res := min(list, len(list))
	return res
}

func main() {
	nums1 := []int{1, 2, 3, 6}
	nums2 := []int{2, 3, 4, 5}
	list := []int{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				list = append(list, nums1[i])
			}
		}
	}
	res := getCommon(nums1, nums2)
	fmt.Println(res)
}
