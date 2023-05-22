package main

import (
	"fmt"
	"regexp"
)

func maxLengthBetweenEqualCharacters(s string) int {
	list := []string{}
	count := 0
	index := []int{}
	start, end := 0, 0
	match, _ := regexp.MatchString("a", s)
	if match == false {
		return -1
	} else {
		for i := 0; i < len(s); i++ {
			char := string(s[i])
			list = append(list, char)
			if char == "a" {
				index = append(index, i)
			}
		}
		start = index[0]
		end = index[1]
	}

	count = len(list[start+1 : end])
	return count

}

// TODO
func main() {
	s := "cabbac"
	res := maxLengthBetweenEqualCharacters(s)
	fmt.Println(res)

	// list := []string{}
	// count := 0
	// index := []int{}
	// start, end := 0, 0
	// match, _ := regexp.MatchString("a", s)
	// if match == false {
	// 	fmt.Println(-1)
	// } else {
	// 	for i := 0; i < len(s); i++ {
	// 		char := string(s[i])
	// 		list = append(list, char)
	// 		if char == "a" {
	// 			index = append(index, i)
	// 		}
	// 	}
	// 	start = index[0]
	// 	end = index[1]
	// }

	// count = len(list[start+1 : end])
	// fmt.Println(count)

}
