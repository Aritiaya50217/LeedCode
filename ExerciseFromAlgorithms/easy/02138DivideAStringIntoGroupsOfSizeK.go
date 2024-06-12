package main

import (
	"fmt"
)

func divideString(s string, k int, fill byte) []string {
	return nil
}

// TODO
func main() {
	s := "abcdefghij"
	k := 3
	fill := "x"
	fmt.Println(len(s), k, fill)
	if len(s)%3 != 0 {
		for len(s)%3 >= 1 {
			s += fill
		}
	}
	fmt.Println(s)

}
