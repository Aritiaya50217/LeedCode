package main

import "fmt"

// res : https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/solutions/1080009/golang-o-1-solution-faster-than-100-with-explanation-and-images/
func countOdds(low int, high int) int {
	return (high+1)/2 - low/2
}

func main() {
	low, high := 8, 10
	res := countOdds(low, high)
	fmt.Println(res)
}
