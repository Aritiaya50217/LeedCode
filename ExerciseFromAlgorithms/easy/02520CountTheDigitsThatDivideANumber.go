package main

import (
	"fmt"
	"strconv"
)

func countDigits(num int) int {
	return 0
}

// TODO
func main() {
	num := 121
	n := strconv.Itoa(num)
	// fmt.Printf("%T : %v\n", n, n)
	for _, v := range n {
		fmt.Println(num % int(v))
	}
}
