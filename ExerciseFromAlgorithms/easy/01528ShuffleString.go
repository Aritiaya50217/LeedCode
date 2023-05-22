package main

import (
	"fmt"
)

// res : https://leetcode.com/problems/shuffle-string/solutions/763747/solutions-in-multiple-languages-java-c-golang-python-kotlin-php/
func restoreString(s string, indices []int) string {
	length := len(s)
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[indices[i]] = s[i]
	}
	return string(res)
}

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	res := restoreString(s, indices)
	fmt.Println(res)
}
