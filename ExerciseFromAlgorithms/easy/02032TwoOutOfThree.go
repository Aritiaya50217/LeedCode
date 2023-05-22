package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	return nil
}

//TODO
func main() {
	nums1 := []int{3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{1, 2}
	fmt.Println(nums1, nums2, nums3)

	list := []int{}
	for _, v := range nums1 {
		for i := 0; i < len(nums2); i++ {
			if v == nums2[i] {
				list = append(list, v, nums2[i])
			}
		}
		for _, val := range nums3 {
			if v == val {
				list = append(list, val)
			}
		}
	}
	fmt.Println(list)

}
