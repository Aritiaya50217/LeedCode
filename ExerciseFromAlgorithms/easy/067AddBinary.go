package main

import "fmt"

// ref : https://leetcode.com/problems/add-binary/solutions/441283/simple-and-elegant-golang-solution-0-ms-faster-than-100-00/?orderBy=most_votes&languageTags=golang
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		// swap
		a, b = b, a
	}
	ans := make([]byte, len(a)+1)
	var carry byte
	for i, j := len(a), len(b); i >= 1 || j >= 1; {
		i, j = i-1, j-1
		var a2Digit, b2Digit byte
		if i >= 0 {
			a2Digit = a[i] - '0'
		}
		if j >= 0 {
			b2Digit = b[j] - '0'
		}
		// sum and carry of full adder
		sum := a2Digit ^ b2Digit ^ carry
		carry = a2Digit&b2Digit | carry&(a2Digit^b2Digit)
		ans[i+1] = sum + '0'
	}
	if carry == 1 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}
func main() {
	a, b := "", ""
	fmt.Scan(&a, &b)
	res := addBinary(a, b)
	fmt.Println(res)
}
