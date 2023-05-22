package main

import (
	"fmt"
	"strconv"
)

func sumBase(n int, k int) int {
	strNum := strconv.FormatInt(int64(n), k)
	list := []int{}
	sum := 0
	for _, v := range strNum {
		number, _ := strconv.Atoi(string(v))
		list = append(list, number)
	}
	for _, val := range list {
		sum += val
	}
	return sum
}
func main() {
	n, k := 34, 6
	res := sumBase(n, k)
	fmt.Println(res)
}
