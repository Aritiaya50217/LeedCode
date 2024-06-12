package main

import (
	"fmt"
	"strconv"
	"strings"
)

func separateDigits(nums []int) []int {
	str := ""
	for _, v := range nums {
		s := strconv.Itoa(v)
		str += s
	}
	re := strings.Split(str, "")
	nums = []int{}
	for i := 0; i < len(re); i++ {
		number, _ := strconv.Atoi(re[i])
		nums = append(nums , number)
	}

	return nums
}

func main() {
	nums := []int{13, 25, 83, 77}
	// str := ""
	// for _, v := range nums {
	// 	s := strconv.Itoa(v)
	// 	str += s
	// }
	// re := strings.Split(str, "")
	// res := []int{}
	// for i := 0; i < len(re); i++ {
	// 	number, _ := strconv.Atoi(re[i])
	// 	res = append(res, number)
	// }
	// fmt.Println(res)

	res := separateDigits(nums)
	fmt.Println(res)

}
