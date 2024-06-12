package main

import (
	"fmt"
	"strconv"
	"strings"
)

func differenceOfSum(nums []int) int {
	sum1, sum2, total := 0, 0, 0
	spList := ""

	for i := 0; i < len(nums); i++ {
		sum1 += nums[i]
		s := strconv.Itoa(nums[i])
		spList += s
	}
	sp := strings.Split(spList, "")
	for j := 0; j < len(sp); j++ {
		n, _ := strconv.Atoi(sp[j])
		sum2 += n
	}
	total = sum1 - sum2

	return total
}
func main() {
	nums := []int{1, 15, 6, 3}
	res := differenceOfSum(nums)
	fmt.Println(res)

}
