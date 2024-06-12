package main

import (
	"fmt"
	"strings"
)

// https://www.geeksforgeeks.org/check-if-the-string-starts-with-specified-prefix-in-golang/
func prefixCount(words []string, pref string) int {
	count := 0
	for _, v := range words {
		check := strings.HasPrefix(v, pref)
		if check == true {
			count++
		}
	}
	return count
}
func main() {
	words := []string{"pay", "attention", "practice", "attend"}
	pref := "at"
	res := prefixCount(words , pref)
	fmt.Println(res)
}
