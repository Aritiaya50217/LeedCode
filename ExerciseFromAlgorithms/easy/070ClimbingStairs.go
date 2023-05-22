package main

import "fmt"

// ref : https://leetcode.com/problems/climbing-stairs/solutions/2810612/simple-golang-solution-with-explanation-beats-100-time-and-memory-no-memoization/
func climbStairs(n int) int {
	res := 0
	last := 0
	secondLast := 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			res = 1
		} else if i == 2 {
			res = 2
		} else {
			res = last + secondLast
		}
		secondLast = last
		last = res
	}
	return res
}

func main() {
	input := 0
	fmt.Scan(&input)

	res := climbStairs(input)
	fmt.Println(res)
}
