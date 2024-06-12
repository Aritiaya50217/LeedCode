package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo(n int) bool {
	res := false
	pow := math.Pow(float64(n), 2)
	if pow == float64(n) {
		res = true
	}
	return res
}

// n = 2^x
func main() {
	n := 16
	pow := math.Pow(float64(n), 2)
	fmt.Println(pow)
	if pow == float64(n) {
		fmt.Println(true)
	}
	fmt.Println(false)
}
