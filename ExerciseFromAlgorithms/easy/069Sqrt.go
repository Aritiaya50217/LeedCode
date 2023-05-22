package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	var y float64 = float64(x)
	res := math.Sqrt(y)
	return int(res)
}

func main() {
	input := 0
	fmt.Scan(&input)
	res := mySqrt(input)
	fmt.Println(res)
}
