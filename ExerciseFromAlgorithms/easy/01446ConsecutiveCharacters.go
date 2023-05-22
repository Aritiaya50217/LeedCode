package main

import "fmt"

func maxPower(s string) int {
	count := 0
	str := "e"
	for _, v := range s {
		if string(v) == str {
			count++
		}
	}
	return count
}
func main() {
	input := "abbcccddddeeeeedcba"
	res := maxPower(input)
	fmt.Println(res)
}
