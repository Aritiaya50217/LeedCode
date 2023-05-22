package main

import (
	"fmt"
	"regexp"
)

func maxRepeating(sequence string, word string) int {
	re, _ := regexp.MatchString(word, sequence)
	count := 0
	if re == false {
		return 0
	} else {
		match, _ := regexp.Compile(word)
		list := match.FindAllString(sequence, -1)
		count = len(list)
	}
	return count
}
// TODO
func main() {
	sequence := "aaabaaaabaaabaaaabaaaabaaaabaaaaba"
	word := "aaaba"
	res := maxRepeating(sequence, word)
	fmt.Println(res)
}
