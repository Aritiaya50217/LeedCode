package main

import (
	"fmt"
	"math"
)

func xorOperation(n int, start int) int {
	num := 0
	list := []int{}
	for i := start; i < n; i++ {
		num = i * 2
		list = append(list, num)
	}

	binary := 0
	for _, v := range list {
		binary = v
	}
	index := 0
	decimalNum, remainder := 0, 0
	for binary != 0 {
		remainder = binary % 10
		binary = binary / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}

	return decimalNum
}

// TODO
func main() {
	n := 3
	start := 4 // [3, 5, 7, 9]
	num := 0
	list := []int{}
	for i := start; i <= (start + n); i++ {
		num = i + 2
		list = append(list, num)
	}
	fmt.Println(list)

	// binary := 0
	// for _, v := range list {
	// 	binary = v
	// }
	// index := 0
	// decimalNum, remainder := 0, 0
	// for binary != 0 {
	// 	remainder = binary % 10
	// 	binary = binary / 10
	// 	decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
	// 	index++
	// }
	// fmt.Println(decimalNum)
	// res := xorOperation(n, start)
	// fmt.Println(res)
}
