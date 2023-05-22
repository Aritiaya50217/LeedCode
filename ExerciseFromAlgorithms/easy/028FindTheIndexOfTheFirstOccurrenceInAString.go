package main

import (
	"fmt"
	"strings"
)

// ref : https://golangbyexample.com/index-first-instance-substring-golang/
func strStr(haystack string, needle string) int {
	res := strings.Index(haystack, needle)
	return res
}

func strStrUseAlgorithms(haystack, needle string) int {
	m := len(needle)
	n := len(haystack)
	for windowStart := 0; windowStart <= n-m; windowStart++ {
		for i := 0; i < m; i++ {
			if needle[i] != haystack[windowStart+i] {
				break
			}
			if i == m-1 {
				return windowStart
			}
		}
	}
	return -1
}

func main() {
	haystack := ""
	needle := ""
	fmt.Scan(&haystack, &needle)
	res := strStr(haystack, needle)
	fmt.Println(res)

	fmt.Println("----- Use Algorithms ------")
	res2 := strStrUseAlgorithms(haystack, needle)
	fmt.Println(res2)

}
