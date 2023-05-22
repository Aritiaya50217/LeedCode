package main

import "fmt"

// ref : https://leetcode.com/problems/counting-bits/solutions/1809099/go-shortest-o-n-solution-with-explanation/
func countBits(n int) []int {
	ones := make([]int, n+1)
	for x := 0; x <= n; x++ {
		// ref : https://sites.google.com/site/khasangkhwbkhumporkaermjava/taw-danein-kar-operators/2
		// >> หมายถึงการเลื่อนบิตไปทางขวามือ (Right Shift) ซึ่งมีค่าเท่ากับ base* (2^(shift))
		ones[x] = ones[x>>1] + x&1
	}
	return ones
}

func main() {
	test := countBits(5)
	fmt.Println(test)

	n := 5
	ones := make([]int, n+1)
	fmt.Println("ones :", ones)
	for x := 0; x <= n; x++ {
		ones[x] = ones[x>>1] + x&1
		fmt.Print("ones[x] : ", ones[x], "\t")
		fmt.Print("ones[x>>1] : ", ones[x>>1], "\t")
		fmt.Println("x&1 : ", x&1)
	}
	fmt.Println(ones)

}

/*
					128		64		32		16		8		4		2		1
ex  5 >> 2			0		0		0		0		0		1		0		1

ex  5 & 3
	5				0		0		0		0		0		1		0		1
	3				0		0		0		0		0		0		1		1
	ans = 7			0		0		0		0		0		1		1		1
*/
