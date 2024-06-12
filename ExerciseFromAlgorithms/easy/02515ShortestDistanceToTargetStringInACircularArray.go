package main

import "fmt"

func closetTarget(words []string, target string, startIndex int) int {
	count := 0
	for _, v := range words[startIndex:] {
		if v == target {
			count++
		}
	}
	if count == 0 {
		count = -1
	}
	return count
}

// TODO
func main() {
	words := []string{"hello", "i", "am", "leetcode", "hello"}
	target := "hello"
	startIndex := 1
	res := closetTarget(words, target, startIndex)
	fmt.Println(res)
}
