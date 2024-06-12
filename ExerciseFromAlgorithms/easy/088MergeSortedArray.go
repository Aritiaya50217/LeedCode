package main

import (
	"fmt"
	"sort"
)

// ref : https://gobyexample.com/sorting
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	slice := append(nums1[:m], nums2[:n]...)
	sort.Ints(slice)
	return slice

}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	res := merge(nums1, m, nums2, n)
	fmt.Println(res)

}
