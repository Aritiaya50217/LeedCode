package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Algorithm => current + x = target

	nums   [4	-2	 5	 0	 6	 3	 2	 7]
	target [0	 1	 2   3   4   5   6   7]

	ex  เริ่มที่ 4                             map
	x = 1 - current						 4 => 0
	x = 1 - 4 = -3

	ex ถึง nums = 3						  3 => 5
	x = 1 - 3 = -2

	ans => 3+(-2) = 1 => o(n)
	target => [5,1]

*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, idx}
		}
		m[num] = idx
	}
	return nil
}

func main() {
	n := bufio.NewScanner(os.Stdin)
	nums := []int{}
	for n.Scan() {
		nStr := n.Text()
		number, _ := strconv.Atoi(nStr)
		if nStr != "" {
			nums = append(nums, number)
		} else {
			break
		}
	}

	target := nums[len(nums)-1:]
	res := twoSum(nums, target[0])
	fmt.Println(res)

}
