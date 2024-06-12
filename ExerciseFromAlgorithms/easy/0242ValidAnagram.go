package main

import (
	"fmt"
	"strings"
)

func isAnagram(s string, t string) bool {
	check := true
	wordOne := strings.Split(s, "")
	wordTwo := make([]string, len(wordOne))
	wordTwo = append(wordTwo, t)
	count := 0
	for i, v := range wordOne {
		count = strings.Count(wordTwo[i], v)
		if count == 0 {
			check = false
		}
	}
	return check
}

// TODO
func main() {
	s, t := "ab", "a"
	wordOne := strings.Split(s, "")
	count := 0
	for _, v := range wordOne {
		count = strings.Count(t, v)
		if count == 0 {
			fmt.Println(false)
		}
	}

	// res := isAnagram(s, t)
	// fmt.Println(res)
}
