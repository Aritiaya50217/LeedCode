package main

import (
	"fmt"
	"strconv"
)

// ref : https://www.tutorialspoint.com/write-a-golang-program-to-reverse-a-number
func reverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}

func isSameAfterReversals(num int) bool {
	resverse := reverseNumber(num)
	numStr := strconv.Itoa(num)
	resverseStr := strconv.Itoa(resverse)
	if len(resverseStr) == len(numStr) {
		return true
	} else {
		return false
	}
	return false
}

func main() {
	num := 1800
	res := isSameAfterReversals(num)
	fmt.Println(res)

}
